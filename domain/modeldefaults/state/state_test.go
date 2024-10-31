// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"database/sql"

	"github.com/canonical/sqlair"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/cloud"
	corecloud "github.com/juju/juju/core/cloud"
	cloudtesting "github.com/juju/juju/core/cloud/testing"
	"github.com/juju/juju/core/model"
	modeltesting "github.com/juju/juju/core/model/testing"
	usertesting "github.com/juju/juju/core/user/testing"
	clouderrors "github.com/juju/juju/domain/cloud/errors"
	cloudstate "github.com/juju/juju/domain/cloud/state"
	modelerrors "github.com/juju/juju/domain/model/errors"
	modelstatetesting "github.com/juju/juju/domain/model/state/testing"
	schematesting "github.com/juju/juju/domain/schema/testing"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/internal/uuid"
)

type stateSuite struct {
	schematesting.ControllerSuite

	modelUUID            model.UUID
	modelCloudUUID       corecloud.UUID
	modelCloudName       string
	modelCloudRegionName string
}

var _ = gc.Suite(&stateSuite{})

func (s *stateSuite) SetUpTest(c *gc.C) {
	s.ControllerSuite.SetUpTest(c)
	s.modelUUID = modelstatetesting.CreateTestModel(c, s.TxnRunnerFactory(), "model-defaults")

	var cloudUUIDStr, cloudName, cloudRegionName string
	err := s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		stmt := "SELECT cloud_uuid, cloud_name, cloud_region_name FROM v_model WHERE uuid = ?"
		err := tx.QueryRowContext(ctx, stmt, s.modelUUID).Scan(&cloudUUIDStr, &cloudName, &cloudRegionName)
		if err != nil {
			return err
		}
		return nil
	})
	c.Assert(err, jc.ErrorIsNil)
	s.modelCloudUUID = corecloud.UUID(cloudUUIDStr)
	s.modelCloudName = cloudName
	s.modelCloudRegionName = cloudRegionName
}

// TestModelMetadataDefaults is asserting the happy path of model metadata
// defaults.
func (s *stateSuite) TestModelMetadataDefaults(c *gc.C) {
	uuid := modelstatetesting.CreateTestModel(c, s.TxnRunnerFactory(), "test")
	st := NewState(s.TxnRunnerFactory())
	defaults, err := st.ModelMetadataDefaults(context.Background(), uuid)
	c.Check(err, jc.ErrorIsNil)
	c.Check(defaults, jc.DeepEquals, map[string]string{
		config.NameKey: "test",
		config.UUIDKey: uuid.String(),
		config.TypeKey: "ec2",
	})
}

// TestModelMetadataDefaultsNoModel is asserting that if we ask for the model
// metadata defaults for a model that doesn't exist we get back a
// [modelerrors.NotFound] error.
func (s *stateSuite) TestModelMetadataDefaultsNoModel(c *gc.C) {
	uuid := modeltesting.GenModelUUID(c)
	st := NewState(s.TxnRunnerFactory())
	defaults, err := st.ModelMetadataDefaults(context.Background(), uuid)
	c.Check(err, jc.ErrorIs, modelerrors.NotFound)
	c.Check(len(defaults), gc.Equals, 0)
}

var (
	testCloud = cloud.Cloud{
		Name:      "fluffy",
		Type:      "ec2",
		AuthTypes: []cloud.AuthType{cloud.AccessKeyAuthType, cloud.UserPassAuthType},
		Endpoint:  "https://endpoint",
		Regions: []cloud.Region{{
			Name: "region1",
		}, {
			Name: "region2",
		}},
	}
)

// TestUpdateCloudDefaults is testing and ensuring that for the simple happy
// case of updating a given cloud's defaults we can both set a set of values and
// then read them back verbatim with the cloud's UUID.
func (s *stateSuite) TestUpdateCloudDefaults(c *gc.C) {
	cloudSt := cloudstate.NewState(s.TxnRunnerFactory())
	cloudUUID := cloudtesting.GenCloudUUID(c)
	err := cloudSt.CreateCloud(context.Background(), usertesting.GenNewName(c, "admin"), cloudUUID.String(), testCloud)
	c.Assert(err, jc.ErrorIsNil)

	st := NewState(s.TxnRunnerFactory())
	err = st.UpdateCloudDefaults(context.Background(), cloudUUID, map[string]string{
		"foo":        "bar",
		"wallyworld": "peachy",
	})
	c.Assert(err, jc.ErrorIsNil)

	defaults, err := st.CloudDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(defaults, jc.DeepEquals, map[string]string{
		"foo":        "bar",
		"wallyworld": "peachy",
	})
}

// TestComplexUpdateCloudDefaults is testing a more complex update strategy for
// a cloud defaults where we perform several overwrite overwrite actions
// ("updates") for a key and also delete another key. At the end we check that
// the reported cloud defaults match the set of updates.
func (s *stateSuite) TestComplexUpdateCloudDefaults(c *gc.C) {
	cloudSt := cloudstate.NewState(s.TxnRunnerFactory())
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	err := cloudSt.CreateCloud(context.Background(), usertesting.GenNewName(c, "admin"), cloudUUID.String(), testCloud)
	c.Assert(err, jc.ErrorIsNil)

	st := NewState(s.TxnRunnerFactory())
	err = st.UpdateCloudDefaults(context.Background(), cloudUUID, map[string]string{
		"foo":        "bar",
		"wallyworld": "peachy",
	})
	c.Assert(err, jc.ErrorIsNil)

	defaults, err := st.CloudDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(defaults, jc.DeepEquals, map[string]string{
		"foo":        "bar",
		"wallyworld": "peachy",
	})

	err = st.UpdateCloudDefaults(context.Background(), cloudUUID, map[string]string{
		"wallyworld": "peachy1",
		"foo2":       "bar2",
	})
	c.Assert(err, jc.ErrorIsNil)

	err = st.DeleteCloudDefaults(context.Background(), cloudUUID, []string{"foo"})
	c.Assert(err, jc.ErrorIsNil)

	defaults, err = st.CloudDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(defaults, jc.DeepEquals, map[string]string{
		"wallyworld": "peachy1",
		"foo2":       "bar2",
	})
}

// TestUpdateCloudDefaultsCloudNotFound is asserting that if we try and update
// default config values for a cloud that does not exist we get back an error
// that satisfies [clouderrors.NotFound]
//
// We also perform the same check but with an empty input set of updates. That
// is because the contract we have that says regardless of the update if the
// cloud does not exist we will still get back an error satisfying
// [clouderrors.NotFound].
func (s *stateSuite) TestUpdateCloudDefaultsCloudNotFound(c *gc.C) {
	cloudUUID := cloudtesting.GenCloudUUID(c)
	err := NewState(s.TxnRunnerFactory()).UpdateCloudDefaults(
		context.Background(),
		cloudUUID,
		map[string]string{
			"foo": "bar",
		},
	)
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}

func (s *stateSuite) TestCloudDefaultsUpdateForNonExistentCloud(c *gc.C) {
	cloudUUID := cloudtesting.GenCloudUUID(c)
	st := NewState(s.TxnRunnerFactory())
	err := st.UpdateCloudDefaults(context.Background(), cloudUUID, map[string]string{
		"wallyworld": "peachy",
	})
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}

// TestUpdateNonExistentCloudRegionDefaults is asserting that if we attempt to
// update the defaults for a cloud region that doesn't exist we get back an
// error satisfying [clouderrors.NotFound].
func (s *stateSuite) TestUpdateNonExistentCloudRegionDefaults(c *gc.C) {
	cloudUUID := cloudtesting.GenCloudUUID(c)
	st := NewState(s.TxnRunnerFactory())
	err := st.UpdateCloudRegionDefaults(
		context.Background(),
		cloudUUID,
		"noexist",
		nil,
	)
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}

// TestCloudDefaultsCloudNotFound is asserting that if we ask for the defaults
// of cloud that doesn't exist we get back an error satisfying
// [clouderrors.NotFound].
func (s *stateSuite) TestCloudDefaultsCloudNotFound(c *gc.C) {
	cloudUUID := cloudtesting.GenCloudUUID(c)
	_, err := NewState(s.TxnRunnerFactory()).CloudDefaults(context.Background(), cloudUUID)
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}

func (s *stateSuite) TestCloudAllRegionDefaults(c *gc.C) {
	cld := testCloud

	cloudSt := cloudstate.NewState(s.TxnRunnerFactory())
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	err := cloudSt.CreateCloud(context.Background(), usertesting.GenNewName(c, "admin"), cloudUUID.String(), cld)
	c.Assert(err, jc.ErrorIsNil)

	st := NewState(s.TxnRunnerFactory())
	err = st.UpdateCloudRegionDefaults(
		context.Background(),
		cloudUUID,
		cld.Regions[0].Name,
		map[string]string{
			"foo":        "bar",
			"wallyworld": "peachy",
		})
	c.Assert(err, jc.ErrorIsNil)

	err = st.UpdateCloudRegionDefaults(
		context.Background(),
		cloudUUID,
		cld.Regions[1].Name,
		map[string]string{
			"foo":        "bar1",
			"wallyworld": "peachy2",
		})
	c.Assert(err, jc.ErrorIsNil)

	regionDefaults, err := st.CloudAllRegionDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(regionDefaults, jc.DeepEquals, map[string]map[string]string{
		cld.Regions[0].Name: {
			"foo":        "bar",
			"wallyworld": "peachy",
		},
		cld.Regions[1].Name: {
			"foo":        "bar1",
			"wallyworld": "peachy2",
		},
	})
}

func (s *stateSuite) TestCloudAllRegionDefaultsComplex(c *gc.C) {
	cld := testCloud

	cloudSt := cloudstate.NewState(s.TxnRunnerFactory())
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	err := cloudSt.CreateCloud(context.Background(), usertesting.GenNewName(c, "admin"), cloudUUID.String(), cld)
	c.Assert(err, jc.ErrorIsNil)

	st := NewState(s.TxnRunnerFactory())
	err = st.UpdateCloudRegionDefaults(
		context.Background(),
		cloudUUID,
		cld.Regions[0].Name,
		map[string]string{
			"foo":        "bar",
			"wallyworld": "peachy",
		})
	c.Assert(err, jc.ErrorIsNil)

	err = st.UpdateCloudRegionDefaults(
		context.Background(),
		cloudUUID,
		cld.Regions[1].Name,
		map[string]string{
			"foo":        "bar1",
			"wallyworld": "peachy2",
		})
	c.Assert(err, jc.ErrorIsNil)

	regionDefaults, err := st.CloudAllRegionDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(regionDefaults, jc.DeepEquals, map[string]map[string]string{
		cld.Regions[0].Name: {
			"foo":        "bar",
			"wallyworld": "peachy",
		},
		cld.Regions[1].Name: {
			"foo":        "bar1",
			"wallyworld": "peachy2",
		},
	})

	err = st.UpdateCloudRegionDefaults(
		context.Background(),
		cloudUUID,
		cld.Regions[1].Name,
		map[string]string{
			"wallyworld": "peachy3",
		})
	c.Assert(err, jc.ErrorIsNil)

	err = st.DeleteCloudRegionDefaults(
		context.Background(),
		cloudUUID,
		cld.Regions[1].Name,
		[]string{"foo"})
	c.Assert(err, jc.ErrorIsNil)

	err = st.UpdateCloudRegionDefaults(
		context.Background(),
		cloudUUID,
		cld.Regions[0].Name,
		map[string]string{
			"one":   "two",
			"three": "four",
			"five":  "six",
		})
	c.Assert(err, jc.ErrorIsNil)

	regionDefaults, err = st.CloudAllRegionDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(regionDefaults, jc.DeepEquals, map[string]map[string]string{
		cld.Regions[0].Name: {
			"foo":        "bar",
			"wallyworld": "peachy",
			"one":        "two",
			"three":      "four",
			"five":       "six",
		},
		cld.Regions[1].Name: {
			"wallyworld": "peachy3",
		},
	})
}

// TestCloudAllRegionDefaultsNoExist is testing that if there are no cloud
// region defaults set for a given cloud an empty map is returned and no errors
// are produced.
func (s *stateSuite) TestCloudAllRegionDefaultsNoExist(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	defaults, err := st.CloudAllRegionDefaults(context.Background(), s.modelCloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(len(defaults), gc.Equals, 0)
}

// TestCloudAllRegionDefaultsCloudNotFound is asserting that if we ask for the
// defaults of every region on a cloud that does not exist we get back an error
// satisfying [clouderrors.NotFound].
func (s *stateSuite) TestCloudAllRegionDefaultsCloudNotFound(c *gc.C) {
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	_, err := NewState(s.TxnRunnerFactory()).CloudAllRegionDefaults(context.Background(), cloudUUID)
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}

func (s *stateSuite) TestModelCloudRegionDefaults(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	err := st.UpdateCloudRegionDefaults(
		context.Background(),
		s.modelCloudUUID,
		s.modelCloudRegionName,
		map[string]string{
			"foo":        "bar",
			"wallyworld": "peachy",
		})
	c.Assert(err, jc.ErrorIsNil)

	regionDefaults, err := st.ModelCloudRegionDefaults(context.Background(), s.modelUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(regionDefaults, jc.DeepEquals, map[string]string{
		"foo":        "bar",
		"wallyworld": "peachy",
	})
}

func (s *stateSuite) TestModelCloudRegionDefaultsNone(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	regionDefaults, err := st.ModelCloudRegionDefaults(context.Background(), s.modelUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(regionDefaults, gc.HasLen, 0)
}

// TestModelCloudRegionDefaults is asserting that if we ask for the cloud region
// defaults for a models cloud region and the model does not exist we get back a
// [modelerrors.NotFound] error.
func (s *stateSuite) TestModelCloudRegionDefaultsNoModel(c *gc.C) {
	uuid := modeltesting.GenModelUUID(c)
	st := NewState(s.TxnRunnerFactory())
	defaults, err := st.ModelCloudRegionDefaults(context.Background(), uuid)
	c.Check(err, jc.ErrorIs, modelerrors.NotFound)
	c.Check(len(defaults), gc.Equals, 0)
}

func (s *stateSuite) TestCloudDefaultsRemoval(c *gc.C) {
	cloudSt := cloudstate.NewState(s.TxnRunnerFactory())
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	err := cloudSt.CreateCloud(context.Background(), usertesting.GenNewName(c, "admin"), cloudUUID.String(), testCloud)
	c.Assert(err, jc.ErrorIsNil)

	st := NewState(s.TxnRunnerFactory())
	err = st.UpdateCloudDefaults(context.Background(), cloudUUID, map[string]string{
		"foo":        "bar",
		"wallyworld": "peachy",
	})
	c.Assert(err, jc.ErrorIsNil)

	err = st.DeleteCloudDefaults(context.Background(), cloudUUID, []string{"foo"})
	c.Assert(err, jc.ErrorIsNil)

	defaults, err := st.CloudDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(defaults, jc.DeepEquals, map[string]string{
		"wallyworld": "peachy",
	})

	err = st.DeleteCloudDefaults(context.Background(), cloudUUID, []string{"noexist"})
	c.Assert(err, jc.ErrorIsNil)

	defaults, err = st.CloudDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(defaults, jc.DeepEquals, map[string]string{
		"wallyworld": "peachy",
	})
}

// TestDeleteCloudDefaultsCloudNotFound is asserting that if we try and delete
// defaults for a cloud that doesn't exist we get back an error satisfying
// [clouderrors.NotFound]
func (s *stateSuite) TestDeleteCloudDefaultsCloudNotFound(c *gc.C) {
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	err := NewState(s.TxnRunnerFactory()).DeleteCloudDefaults(context.Background(), cloudUUID, []string{"foo"})
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}

func (s *stateSuite) TestEmptyCloudDefaults(c *gc.C) {
	cloudSt := cloudstate.NewState(s.TxnRunnerFactory())
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	err := cloudSt.CreateCloud(context.Background(), usertesting.GenNewName(c, "admin"), cloudUUID.String(), testCloud)
	c.Assert(err, jc.ErrorIsNil)

	st := NewState(s.TxnRunnerFactory())
	defaults, err := st.CloudDefaults(context.Background(), cloudUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(len(defaults), gc.Equals, 0)
}

func (s *stateSuite) TestGetCloudUUID(c *gc.C) {
	cloudSt := cloudstate.NewState(s.TxnRunnerFactory())
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	err := cloudSt.CreateCloud(context.Background(), usertesting.GenNewName(c, "admin"), cloudUUID.String(), testCloud)
	c.Assert(err, jc.ErrorIsNil)

	st := NewState(s.TxnRunnerFactory())
	uuid, err := st.GetCloudUUID(context.Background(), testCloud.Name)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(uuid.String(), gc.Equals, cloudUUID.String())
}

// TestGetCloudUUIDNotFound is asserting that if we ask for the UUID of a cloud
// name that does not exist we get back an error satisfying
// [clouderrors.NotFound].
func (s *stateSuite) TestGetCloudUUIDNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	_, err := st.GetCloudUUID(context.Background(), "noexist")
	c.Assert(err, jc.ErrorIs, clouderrors.NotFound)
}

func (s *stateSuite) TestGetModelCloudUUID(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	gotCloudUUID, err := st.GetModelCloudUUID(context.Background(), s.modelUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(gotCloudUUID.String(), gc.Equals, s.modelCloudUUID.String())
}

// TestGetModelCloudType asserts that the cloud type for a created model is
// correct.
func (s *stateSuite) TestGetCloudType(c *gc.C) {
	cloudSt := cloudstate.NewState(s.TxnRunnerFactory())
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	err := cloudSt.CreateCloud(context.Background(), usertesting.GenNewName(c, "admin"), cloudUUID.String(), testCloud)
	c.Assert(err, jc.ErrorIsNil)

	ct, err := NewState(s.TxnRunnerFactory()).CloudType(
		context.Background(), cloudUUID,
	)
	c.Check(err, jc.ErrorIsNil)
	c.Check(ct, gc.Equals, "ec2")
}

// TestGetModelCloudTypModelNotFound is asserting that when no model exists we
// get back a [modelerrors.NotFound] error when querying for a model's cloud
// type.
func (s *stateSuite) TestGetCloudTypeCloudNotFound(c *gc.C) {
	cloudUUID := corecloud.UUID(uuid.MustNewUUID().String())
	_, err := NewState(s.TxnRunnerFactory()).CloudType(
		context.Background(), cloudUUID,
	)
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}

// TestSetCloudDefaults is asserting that if we set cloud defaults for a cloud
// they are reflected back to use when retrieving the cloud defaults.
func (s *stateSuite) TestSetCloudDefaults(c *gc.C) {
	err := s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		return SetCloudDefaults(ctx, tx, s.modelCloudName, map[string]string{
			"foo": "bar",
		})
	})
	c.Check(err, jc.ErrorIsNil)

	defaults, err := NewState(s.TxnRunnerFactory()).CloudDefaults(
		context.Background(), s.modelCloudUUID,
	)
	c.Check(err, jc.ErrorIsNil)
	c.Check(defaults, gc.DeepEquals, map[string]string{
		"foo": "bar",
	})
}

// TestSetCloudDefaultsOverrides is asserting that [SetCloudDefaults] overrides
// any previously set values for a cloud's defaults.
func (s *stateSuite) TestSetCloudDefaultsOverrides(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	err := st.UpdateCloudDefaults(context.Background(), s.modelCloudUUID, map[string]string{
		"testkey": "testval",
	})
	c.Check(err, jc.ErrorIsNil)

	err = s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		return SetCloudDefaults(ctx, tx, s.modelCloudName, map[string]string{
			"foo": "bar",
		})
	})
	c.Check(err, jc.ErrorIsNil)

	defaults, err := st.CloudDefaults(context.Background(), s.modelCloudUUID)
	c.Check(err, jc.ErrorIsNil)
	c.Check(defaults, gc.DeepEquals, map[string]string{
		"foo": "bar",
	})
}

// TestSetCloudDefaultsRemoves is asserting that if we pass an empty set of
// defaults to [SetCloudDefaults] we remove all the defaults for a cloud that
// have already been set.
func (s *stateSuite) TestSetCloudDefaultsRemoves(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	err := st.UpdateCloudDefaults(context.Background(), s.modelCloudUUID, map[string]string{
		"testkey": "testval",
	})
	c.Check(err, jc.ErrorIsNil)

	err = s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		return SetCloudDefaults(ctx, tx, s.modelCloudName, map[string]string{})
	})
	c.Check(err, jc.ErrorIsNil)

	defaults, err := st.CloudDefaults(context.Background(), s.modelCloudUUID)
	c.Check(err, jc.ErrorIsNil)
	c.Check(defaults, gc.HasLen, 0)
}

// TestSetCloudDefaultsCloudNotFound is asserting that if we try and set cloud
// defaults for a cloud that does not exist we get back an error that satisfies
// [clouderrors.NotFound].
func (s *stateSuite) TestSetCloudDefaultsCloudNotFound(c *gc.C) {
	err := s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		return SetCloudDefaults(ctx, tx, "noexist", map[string]string{})
	})
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}

// TestDeleteCloudRegionDefaultsCloudNotFound is testing that we get a an error
// satisfying [clouderrors.NotFound] when we try and delete cloud region
// defaults for either a cloud or region that doesn't exist.
func (s *stateSuite) TestDeleteCloudRegionDefaultsCloudNotFound(c *gc.C) {
	err := NewState(s.TxnRunnerFactory()).DeleteCloudRegionDefaults(
		context.Background(), s.modelCloudUUID, "noexist", []string{"foo"},
	)
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)

	cloudUUID := cloudtesting.GenCloudUUID(c)
	err = NewState(s.TxnRunnerFactory()).DeleteCloudRegionDefaults(
		context.Background(), cloudUUID, "noexist", []string{"foo"},
	)
	c.Check(err, jc.ErrorIs, clouderrors.NotFound)
}
