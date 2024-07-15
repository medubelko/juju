// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import (
	"github.com/juju/errors"
)

const (
	// NameNotValid describes an error that occurs when attempting to get
	// a charm using an invalid name.
	NameNotValid = errors.ConstError("charm name not valid")

	// CharmSourceNotValid describes an error that occurs when attempting to get
	// a charm using an invalid charm source.
	CharmSourceNotValid = errors.ConstError("charm source not valid")

	// NotFound describes an error that occurs when a charm cannot be found.
	NotFound = errors.ConstError("charm not found")

	// AlreadyExists describes an error that occurs when a charm already
	// exists for the given natural key.
	AlreadyExists = errors.ConstError("charm already exists")

	// RevisionNotValid describes an error that occurs when attempting to get
	// a charm using an invalid revision.
	RevisionNotValid = errors.ConstError("charm revision not valid")

	// MetadataNotValid describes an error that occurs when the charm metadata
	// is not valid.
	MetadataNotValid = errors.ConstError("charm metadata not valid")

	// ManifestNotValid describes an error that occurs when the charm manifest
	// is not valid.
	ManifestNotValid = errors.ConstError("charm manifest not valid")
)
