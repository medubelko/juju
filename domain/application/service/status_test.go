// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"time"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/status"
	coreunit "github.com/juju/juju/core/unit"
	"github.com/juju/juju/domain/application"
)

type statusSuite struct{}

var _ = gc.Suite(&statusSuite{})

var now = time.Now()

func (s *statusSuite) TestEncodeCloudContainerStatus(c *gc.C) {
	testCases := []struct {
		input  *status.StatusInfo
		output *application.StatusInfo[application.CloudContainerStatusType]
	}{
		{
			input: &status.StatusInfo{
				Status: status.Waiting,
			},
			output: &application.StatusInfo[application.CloudContainerStatusType]{
				Status: application.CloudContainerStatusWaiting,
			},
		},
		{
			input: &status.StatusInfo{
				Status: status.Blocked,
			},
			output: &application.StatusInfo[application.CloudContainerStatusType]{
				Status: application.CloudContainerStatusBlocked,
			},
		},
		{
			input: &status.StatusInfo{
				Status: status.Running,
			},
			output: &application.StatusInfo[application.CloudContainerStatusType]{
				Status: application.CloudContainerStatusRunning,
			},
		},
		{
			input: &status.StatusInfo{
				Status:  status.Running,
				Message: "I'm active!",
				Data:    map[string]interface{}{"foo": "bar"},
				Since:   &now,
			},
			output: &application.StatusInfo[application.CloudContainerStatusType]{
				Status:  application.CloudContainerStatusRunning,
				Message: "I'm active!",
				Data:    []byte(`{"foo":"bar"}`),
				Since:   &now,
			},
		},
	}

	for i, test := range testCases {
		c.Logf("test %d: %v", i, test.input)
		output, err := encodeCloudContainerStatus(test.input)
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(output, jc.DeepEquals, test.output)
		result, err := decodeCloudContainerStatus(output)
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(result, jc.DeepEquals, test.input)
	}
}

func (s *statusSuite) TestEncodeWorkloadStatus(c *gc.C) {
	testCases := []struct {
		input  *status.StatusInfo
		output *application.StatusInfo[application.WorkloadStatusType]
	}{
		{
			input: &status.StatusInfo{
				Status: status.Unset,
			},
			output: &application.StatusInfo[application.WorkloadStatusType]{
				Status: application.WorkloadStatusUnset,
			},
		},
		{
			input: &status.StatusInfo{
				Status: status.Unknown,
			},
			output: &application.StatusInfo[application.WorkloadStatusType]{
				Status: application.WorkloadStatusUnknown,
			},
		},
		{
			input: &status.StatusInfo{
				Status: status.Maintenance,
			},
			output: &application.StatusInfo[application.WorkloadStatusType]{
				Status: application.WorkloadStatusMaintenance,
			},
		},
		{
			input: &status.StatusInfo{
				Status: status.Waiting,
			},
			output: &application.StatusInfo[application.WorkloadStatusType]{
				Status: application.WorkloadStatusWaiting,
			},
		},
		{
			input: &status.StatusInfo{
				Status: status.Blocked,
			},
			output: &application.StatusInfo[application.WorkloadStatusType]{
				Status: application.WorkloadStatusBlocked,
			},
		},
		{
			input: &status.StatusInfo{
				Status: status.Active,
			},
			output: &application.StatusInfo[application.WorkloadStatusType]{
				Status: application.WorkloadStatusActive,
			},
		},
		{
			input: &status.StatusInfo{
				Status: status.Terminated,
			},
			output: &application.StatusInfo[application.WorkloadStatusType]{
				Status: application.WorkloadStatusTerminated,
			},
		},
		{
			input: &status.StatusInfo{
				Status:  status.Active,
				Message: "I'm active!",
				Data:    map[string]interface{}{"foo": "bar"},
				Since:   &now,
			},
			output: &application.StatusInfo[application.WorkloadStatusType]{
				Status:  application.WorkloadStatusActive,
				Message: "I'm active!",
				Data:    []byte(`{"foo":"bar"}`),
				Since:   &now,
			},
		},
	}

	for i, test := range testCases {
		c.Logf("test %d: %v", i, test.input)
		output, err := encodeWorkloadStatus(test.input)
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(output, jc.DeepEquals, test.output)
		result, err := decodeWorkloadStatus(output)
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(result, jc.DeepEquals, test.input)
	}
}

func (s *statusSuite) TestReduceWorkloadStatusesEmpty(c *gc.C) {
	info, err := reduceWorkloadStatuses(nil)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Unknown,
	})

	info, err = reduceWorkloadStatuses([]application.StatusInfo[application.WorkloadStatusType]{})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Unknown,
	})
}

func (s *statusSuite) TestReduceWorkloadStatusesBringsAllDetails(c *gc.C) {
	value := application.StatusInfo[application.WorkloadStatusType]{
		Status:  application.WorkloadStatusActive,
		Message: "I'm active",
		Data:    []byte(`{"key":"value"}`),
		Since:   &now,
	}
	info, err := reduceWorkloadStatuses([]application.StatusInfo[application.WorkloadStatusType]{value})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status:  status.Active,
		Message: "I'm active",
		Data:    map[string]interface{}{"key": "value"},
		Since:   &now,
	})
}

func (s *statusSuite) TestReduceWorkloadStatusesPriority(c *gc.C) {
	for _, t := range []struct {
		status1  application.WorkloadStatusType
		status2  application.WorkloadStatusType
		expected status.Status
	}{
		// Waiting trumps active
		{application.WorkloadStatusActive, application.WorkloadStatusWaiting, status.Waiting},

		// Maintenance trumps active
		{application.WorkloadStatusMaintenance, application.WorkloadStatusWaiting, status.Maintenance},

		// Blocked trumps active
		{application.WorkloadStatusActive, application.WorkloadStatusBlocked, status.Blocked},

		// Blocked trumps waiting
		{application.WorkloadStatusWaiting, application.WorkloadStatusBlocked, status.Blocked},

		// Blocked trumps maintenance
		{application.WorkloadStatusMaintenance, application.WorkloadStatusBlocked, status.Blocked},
	} {
		value, err := reduceWorkloadStatuses([]application.StatusInfo[application.WorkloadStatusType]{
			{Status: t.status1}, {Status: t.status2},
		})
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(value, gc.NotNil)
		c.Check(value.Status, gc.Equals, t.expected)
	}
}

func (s *statusSuite) TestUnitDisplayStatusNoContainer(c *gc.C) {
	workloadStatus := &application.StatusInfo[application.WorkloadStatusType]{
		Status:  application.WorkloadStatusActive,
		Message: "I'm active",
		Data:    []byte(`{"key":"value"}`),
		Since:   &now,
	}

	info, err := unitDisplayStatus(workloadStatus, nil)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status:  status.Active,
		Message: "I'm active",
		Data:    map[string]interface{}{"key": "value"},
		Since:   &now,
	})
}

func (s *statusSuite) TestUnitDisplayStatusWorkloadTerminatedBlockedMaintainanceDominates(c *gc.C) {
	containerStatus := &application.StatusInfo[application.CloudContainerStatusType]{
		Status: application.CloudContainerStatusBlocked,
	}

	workloadStatus := &application.StatusInfo[application.WorkloadStatusType]{
		Status:  application.WorkloadStatusTerminated,
		Message: "msg",
		Data:    []byte(`{"key":"value"}`),
		Since:   &now,
	}

	expected := &status.StatusInfo{
		Status:  status.Terminated,
		Message: "msg",
		Data:    map[string]interface{}{"key": "value"},
		Since:   &now,
	}

	info, err := unitDisplayStatus(workloadStatus, containerStatus)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, expected)

	workloadStatus.Status = application.WorkloadStatusBlocked
	expected.Status = status.Blocked
	info, err = unitDisplayStatus(workloadStatus, containerStatus)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, expected)

	workloadStatus.Status = application.WorkloadStatusMaintenance
	expected.Status = status.Maintenance
	info, err = unitDisplayStatus(workloadStatus, containerStatus)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, expected)
}

func (s *statusSuite) TestUnitDisplayStatusContainerBlockedDominates(c *gc.C) {
	workloadStatus := &application.StatusInfo[application.WorkloadStatusType]{
		Status: application.WorkloadStatusWaiting,
	}

	containerStatus := &application.StatusInfo[application.CloudContainerStatusType]{
		Status:  application.CloudContainerStatusBlocked,
		Message: "msg",
		Data:    []byte(`{"key":"value"}`),
		Since:   &now,
	}

	info, err := unitDisplayStatus(workloadStatus, containerStatus)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status:  status.Blocked,
		Message: "msg",
		Data:    map[string]interface{}{"key": "value"},
		Since:   &now,
	})
}

func (s *statusSuite) TestUnitDisplayStatusContainerWaitingDominatesActiveWorkload(c *gc.C) {
	workloadStatus := &application.StatusInfo[application.WorkloadStatusType]{
		Status: application.WorkloadStatusActive,
	}

	containerStatus := &application.StatusInfo[application.CloudContainerStatusType]{
		Status:  application.CloudContainerStatusWaiting,
		Message: "msg",
		Data:    []byte(`{"key":"value"}`),
		Since:   &now,
	}

	info, err := unitDisplayStatus(workloadStatus, containerStatus)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status:  status.Waiting,
		Message: "msg",
		Data:    map[string]interface{}{"key": "value"},
		Since:   &now,
	})
}

func (s *statusSuite) TestUnitDisplayStatusContainerRunningDominatesWaitingWorkload(c *gc.C) {
	workloadStatus := &application.StatusInfo[application.WorkloadStatusType]{
		Status: application.WorkloadStatusWaiting,
	}

	containerStatus := &application.StatusInfo[application.CloudContainerStatusType]{
		Status:  application.CloudContainerStatusRunning,
		Message: "msg",
		Data:    []byte(`{"key":"value"}`),
		Since:   &now,
	}

	info, err := unitDisplayStatus(workloadStatus, containerStatus)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status:  status.Running,
		Message: "msg",
		Data:    map[string]interface{}{"key": "value"},
		Since:   &now,
	})
}

func (s *statusSuite) TestUnitDisplayStatusDefaultsToWorkload(c *gc.C) {
	workloadStatus := &application.StatusInfo[application.WorkloadStatusType]{
		Status:  application.WorkloadStatusActive,
		Message: "I'm an active workload",
	}

	containerStatus := &application.StatusInfo[application.CloudContainerStatusType]{
		Status:  application.CloudContainerStatusRunning,
		Message: "I'm a running container",
	}

	info, err := unitDisplayStatus(workloadStatus, containerStatus)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status:  status.Active,
		Message: "I'm an active workload",
	})
}

const (
	unitUUID1 = coreunit.UUID("unit-1")
	unitUUID2 = coreunit.UUID("unit-2")
	unitUUID3 = coreunit.UUID("unit-3")
)

func (s *statusSuite) TestApplicationDisplayStatusFromUnitsNoContainers(c *gc.C) {
	workloadStatuses := map[coreunit.UUID]application.StatusInfo[application.WorkloadStatusType]{
		unitUUID1: {Status: application.WorkloadStatusActive},
		unitUUID2: {Status: application.WorkloadStatusActive},
	}

	info, err := applicationDisplayStatusFromUnits(workloadStatuses, nil)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Active,
	})

	info, err = applicationDisplayStatusFromUnits(
		workloadStatuses,
		make(map[coreunit.UUID]application.StatusInfo[application.CloudContainerStatusType]),
	)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Active,
	})
}

func (s *statusSuite) TestApplicationDisplayStatusFromUnitsEmpty(c *gc.C) {
	info, err := applicationDisplayStatusFromUnits(nil, nil)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Unknown,
	})

	info, err = applicationDisplayStatusFromUnits(
		map[coreunit.UUID]application.StatusInfo[application.WorkloadStatusType]{},
		map[coreunit.UUID]application.StatusInfo[application.CloudContainerStatusType]{},
	)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Unknown,
	})
}

func (s *statusSuite) TestApplicationDisplayStatusFromUnitsPicksGreatestPrecedenceContainer(c *gc.C) {
	workloadStatuses := map[coreunit.UUID]application.StatusInfo[application.WorkloadStatusType]{
		unitUUID1: {Status: application.WorkloadStatusActive},
		unitUUID2: {Status: application.WorkloadStatusActive},
	}

	containerStatuses := map[coreunit.UUID]application.StatusInfo[application.CloudContainerStatusType]{
		unitUUID1: {Status: application.CloudContainerStatusRunning},
		unitUUID2: {Status: application.CloudContainerStatusBlocked},
	}

	info, err := applicationDisplayStatusFromUnits(workloadStatuses, containerStatuses)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Blocked,
	})
}

func (s *statusSuite) TestApplicationDisplayStatusFromUnitsPicksGreatestPrecedenceWorkload(c *gc.C) {
	workloadStatuses := map[coreunit.UUID]application.StatusInfo[application.WorkloadStatusType]{
		unitUUID1: {Status: application.WorkloadStatusActive},
		unitUUID2: {Status: application.WorkloadStatusMaintenance},
	}

	containerStatuses := map[coreunit.UUID]application.StatusInfo[application.CloudContainerStatusType]{
		unitUUID1: {Status: application.CloudContainerStatusRunning},
		unitUUID2: {Status: application.CloudContainerStatusBlocked},
	}

	info, err := applicationDisplayStatusFromUnits(workloadStatuses, containerStatuses)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Maintenance,
	})
}

func (s *statusSuite) TestApplicationDisplayStatusFromUnitsPrioritisesUnitWithGreatestStatusPrecedence(c *gc.C) {
	workloadStatuses := map[coreunit.UUID]application.StatusInfo[application.WorkloadStatusType]{
		unitUUID1: {Status: application.WorkloadStatusActive},
		unitUUID2: {Status: application.WorkloadStatusMaintenance},
	}

	containerStatuses := map[coreunit.UUID]application.StatusInfo[application.CloudContainerStatusType]{
		unitUUID1: {Status: application.CloudContainerStatusBlocked},
		unitUUID2: {Status: application.CloudContainerStatusRunning},
	}

	info, err := applicationDisplayStatusFromUnits(workloadStatuses, containerStatuses)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(info, jc.DeepEquals, &status.StatusInfo{
		Status: status.Blocked,
	})
}
