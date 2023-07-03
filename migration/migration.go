// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package migration

import (
	"context"
	"io"
	"net/url"
	"os"

	"github.com/juju/charm/v11"
	"github.com/juju/description/v4"
	"github.com/juju/errors"
	"github.com/juju/loggo"
	"github.com/juju/naturalsort"
	"github.com/juju/version/v2"

	"github.com/juju/juju/core/migration"
	"github.com/juju/juju/core/modelmigration"
	"github.com/juju/juju/core/resources"
	migrations "github.com/juju/juju/domain/modelmigration"
	"github.com/juju/juju/state"
	"github.com/juju/juju/tools"
)

var logger = loggo.GetLogger("juju.migration")

// StateExporter describes interface on state required to export a
// model.
type StateExporter interface {
	// Export generates an abstract representation of a model.
	Export(leaders map[string]string) (description.Model, error)
}

// StateImporter describes the method needed to import a model
// into the database.
type StateImporter interface {
	Import(model description.Model) (*state.Model, *state.State, error)
}

// ImportModel deserializes a model description from the bytes, transforms
// the model config based on information from the controller model, and then
// imports that as a new database model.
func ImportModel(ctx context.Context, importer StateImporter, scope modelmigration.Scope, bytes []byte) (*state.Model, *state.State, error) {
	model, err := description.Deserialize(bytes)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}

	dbModel, dbState, err := importer.Import(model)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}

	coordinator := modelmigration.NewCoordinator()
	migrations.ImportOperations(coordinator, logger)
	if err := coordinator.Perform(ctx, scope, model); err != nil {
		return nil, nil, errors.Trace(err)
	}

	return dbModel, dbState, nil
}

// CharmDownloader defines a single method that is used to download a
// charm from the source controller in a migration.
type CharmDownloader interface {
	OpenCharm(*charm.URL) (io.ReadCloser, error)
}

// CharmUploader defines a single method that is used to upload a
// charm to the target controller in a migration.
type CharmUploader interface {
	UploadCharm(*charm.URL, io.ReadSeeker) (*charm.URL, error)
}

// ToolsDownloader defines a single method that is used to download
// tools from the source controller in a migration.
type ToolsDownloader interface {
	OpenURI(string, url.Values) (io.ReadCloser, error)
}

// ToolsUploader defines a single method that is used to upload tools
// to the target controller in a migration.
type ToolsUploader interface {
	UploadTools(io.ReadSeeker, version.Binary) (tools.List, error)
}

// ResourceDownloader defines the interface for downloading resources
// from the source controller during a migration.
type ResourceDownloader interface {
	OpenResource(string, string) (io.ReadCloser, error)
}

// ResourceUploader defines the interface for uploading resources into
// the target controller during a migration.
type ResourceUploader interface {
	UploadResource(resources.Resource, io.ReadSeeker) error
	SetPlaceholderResource(resources.Resource) error
	SetUnitResource(string, resources.Resource) error
}

// UploadBinariesConfig provides all the configuration that the
// UploadBinaries function needs to operate. To construct the config
// with the default helper functions, use `NewUploadBinariesConfig`.
type UploadBinariesConfig struct {
	Charms          []string
	CharmDownloader CharmDownloader
	CharmUploader   CharmUploader

	Tools           map[version.Binary]string
	ToolsDownloader ToolsDownloader
	ToolsUploader   ToolsUploader

	Resources          []migration.SerializedModelResource
	ResourceDownloader ResourceDownloader
	ResourceUploader   ResourceUploader
}

// Validate makes sure that all the config values are non-nil.
func (c *UploadBinariesConfig) Validate() error {
	if c.CharmDownloader == nil {
		return errors.NotValidf("missing CharmDownloader")
	}
	if c.CharmUploader == nil {
		return errors.NotValidf("missing CharmUploader")
	}
	if c.ToolsDownloader == nil {
		return errors.NotValidf("missing ToolsDownloader")
	}
	if c.ToolsUploader == nil {
		return errors.NotValidf("missing ToolsUploader")
	}
	if c.ResourceDownloader == nil {
		return errors.NotValidf("missing ResourceDownloader")
	}
	if c.ResourceUploader == nil {
		return errors.NotValidf("missing ResourceUploader")
	}
	return nil
}

// UploadBinaries will send binaries stored in the source blobstore to
// the target controller.
func UploadBinaries(config UploadBinariesConfig) error {
	if err := config.Validate(); err != nil {
		return errors.Trace(err)
	}
	if err := uploadCharms(config); err != nil {
		return errors.Annotatef(err, "cannot upload charms")
	}
	if err := uploadTools(config); err != nil {
		return errors.Annotatef(err, "cannot upload agent binaries")
	}
	if err := uploadResources(config); err != nil {
		return errors.Annotatef(err, "cannot upload resources")
	}
	return nil
}

func streamThroughTempFile(r io.Reader) (_ io.ReadSeeker, cleanup func(), err error) {
	tempFile, err := os.CreateTemp("", "juju-migrate-binary")
	if err != nil {
		return nil, nil, errors.Trace(err)
	}
	defer func() {
		if err != nil {
			_ = os.Remove(tempFile.Name())
		}
	}()
	_, err = io.Copy(tempFile, r)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}
	_, err = tempFile.Seek(0, 0)
	if err != nil {
		return nil, nil, errors.Annotatef(err, "potentially corrupt binary")
	}
	rmTempFile := func() {
		filename := tempFile.Name()
		_ = tempFile.Close()
		_ = os.Remove(filename)
	}

	return tempFile, rmTempFile, nil
}

func uploadCharms(config UploadBinariesConfig) error {
	// It is critical that charms are uploaded in ascending charm URL
	// order so that charm revisions end up the same in the target as
	// they were in the source.
	naturalsort.Sort(config.Charms)

	for _, charmURL := range config.Charms {
		logger.Debugf("sending charm %s to target", charmURL)

		curl, err := charm.ParseURL(charmURL)
		if err != nil {
			return errors.Annotate(err, "bad charm URL")
		}

		reader, err := config.CharmDownloader.OpenCharm(curl)
		if err != nil {
			return errors.Annotate(err, "cannot open charm")
		}
		defer func() { _ = reader.Close() }()

		content, cleanup, err := streamThroughTempFile(reader)
		if err != nil {
			return errors.Trace(err)
		}
		defer cleanup()

		if usedCurl, err := config.CharmUploader.UploadCharm(curl, content); err != nil {
			return errors.Annotate(err, "cannot upload charm")
		} else if usedCurl.String() != curl.String() {
			// The target controller shouldn't assign a different charm URL.
			return errors.Errorf("charm %s unexpectedly assigned %s", curl, usedCurl)
		}
	}
	return nil
}

func uploadTools(config UploadBinariesConfig) error {
	for v, uri := range config.Tools {
		logger.Debugf("sending agent binaries to target: %s", v)

		reader, err := config.ToolsDownloader.OpenURI(uri, nil)
		if err != nil {
			return errors.Annotate(err, "cannot open charm")
		}
		defer func() { _ = reader.Close() }()

		content, cleanup, err := streamThroughTempFile(reader)
		if err != nil {
			return errors.Trace(err)
		}
		defer cleanup()

		if _, err := config.ToolsUploader.UploadTools(content, v); err != nil {
			return errors.Annotate(err, "cannot upload agent binaries")
		}
	}
	return nil
}

func uploadResources(config UploadBinariesConfig) error {
	for _, res := range config.Resources {
		if res.ApplicationRevision.IsPlaceholder() {
			// Resource placeholders created in the migration import rather
			// than attempting to post empty resources.
		} else {
			err := uploadAppResource(config, res.ApplicationRevision)
			if err != nil {
				return errors.Trace(err)
			}
		}
		for unitName, unitRev := range res.UnitRevisions {
			if err := config.ResourceUploader.SetUnitResource(unitName, unitRev); err != nil {
				return errors.Annotate(err, "cannot set unit resource")
			}
		}
		// Each config.Resources element also contains a
		// CharmStoreRevision field. This isn't especially important
		// to migrate so is skipped for now.
	}
	return nil
}

func uploadAppResource(config UploadBinariesConfig, rev resources.Resource) error {
	logger.Debugf("opening application resource for %s: %s", rev.ApplicationID, rev.Name)
	reader, err := config.ResourceDownloader.OpenResource(rev.ApplicationID, rev.Name)
	if err != nil {
		return errors.Annotate(err, "cannot open resource")
	}
	defer func() { _ = reader.Close() }()

	// TODO(menn0) - validate that the downloaded revision matches
	// the expected metadata. Check revision and fingerprint.

	content, cleanup, err := streamThroughTempFile(reader)
	if err != nil {
		return errors.Trace(err)
	}
	defer cleanup()

	if err := config.ResourceUploader.UploadResource(rev, content); err != nil {
		return errors.Annotate(err, "cannot upload resource")
	}
	return nil
}
