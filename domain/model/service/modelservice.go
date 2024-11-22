// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	coremodel "github.com/juju/juju/core/model"
	corestatus "github.com/juju/juju/core/status"
	"github.com/juju/juju/domain/model"
	modelerrors "github.com/juju/juju/domain/model/errors"
	internalerrors "github.com/juju/juju/internal/errors"
	"github.com/juju/juju/internal/uuid"
)

// ModelState is the model state required by this service. This is the model
// database state, not the controller state.
type ModelState interface {
	// Create creates a new model with all of its associated metadata.
	Create(context.Context, model.ReadOnlyModelCreationArgs) error

	// Delete deletes a model.
	Delete(context.Context, coremodel.UUID) error

	// Model returns the read only model information set in the database.
	Model(context.Context) (coremodel.ReadOnlyModel, error)

	// GetStatus returns the status of the model.
	GetStatus(context.Context) (model.StatusInfo, error)

	// SetStatus sets the status of the model.
	SetStatus(context.Context, model.SetStatusArg) error
}

// ControllerState represents the state required for reading all model
// information.
type ControllerState interface {
	// GetModel returns the model with the given UUID.
	GetModel(context.Context, coremodel.UUID) (coremodel.Model, error)

	// GetModelState returns the model state for the given model.
	GetModelState(context.Context, coremodel.UUID) (model.ModelState, error)
}

// ModelService defines a service for interacting with the underlying model
// state, as opposed to the controller state.
type ModelService struct {
	modelID      coremodel.UUID
	controllerSt ControllerState
	modelSt      ModelState
}

// NewModelService returns a new Service for interacting with a models state.
func NewModelService(
	modelID coremodel.UUID,
	controllerSt ControllerState,
	st ModelState,
) *ModelService {
	return &ModelService{
		modelID:      modelID,
		controllerSt: controllerSt,
		modelSt:      st,
	}
}

// GetModelInfo returns the readonly model information for the model in
// question.
func (s *ModelService) GetModelInfo(ctx context.Context) (coremodel.ReadOnlyModel, error) {
	return s.modelSt.Model(ctx)
}

// CreateModel is responsible for creating a new model within the model
// database.
//
// The following error types can be expected to be returned:
// - [modelerrors.AlreadyExists]: When the model uuid is already in use.
func (s *ModelService) CreateModel(
	ctx context.Context,
	controllerUUID uuid.UUID,
) error {
	m, err := s.controllerSt.GetModel(ctx, s.modelID)
	if err != nil {
		return err
	}

	args := model.ReadOnlyModelCreationArgs{
		UUID:            m.UUID,
		AgentVersion:    m.AgentVersion,
		ControllerUUID:  controllerUUID,
		Name:            m.Name,
		Type:            m.ModelType,
		Cloud:           m.Cloud,
		CloudType:       m.CloudType,
		CloudRegion:     m.CloudRegion,
		CredentialOwner: m.Credential.Owner,
		CredentialName:  m.Credential.Name,
	}

	return s.modelSt.Create(ctx, args)
}

// DeleteModel is responsible for removing a model from the system.
//
// The following error types can be expected to be returned:
// - [modelerrors.NotFound]: When the model does not exist.
func (s *ModelService) DeleteModel(
	ctx context.Context,
) error {
	return s.modelSt.Delete(ctx, s.modelID)
}

// Status returns the status of the model.
func (s *ModelService) Status(ctx context.Context) (model.StatusInfo, error) {
	modelState, err := s.controllerSt.GetModelState(ctx, s.modelID)
	if err != nil {
		return model.StatusInfo{}, err
	}
	if modelState.Destroying {
		return model.StatusInfo{
			Status:  corestatus.Destroying,
			Message: "the model is being destroyed",
		}, nil
	}
	if modelState.InvalidCloudCredentialReason != "" {
		return model.StatusInfo{
			Status:  corestatus.Suspended,
			Message: "suspended since cloud credential is not valid",
			Reason:  modelState.InvalidCloudCredentialReason,
		}, nil
	}

	statusInfo, err := s.modelSt.GetStatus(ctx)
	if err != nil {
		return model.StatusInfo{}, internalerrors.Capture(err)
	}
	return statusInfo, nil
}

// validSettableModelStatus returns true if status has a valid value (that is to say,
// a value that it's OK to set) for models.
func validSettableModelStatus(status corestatus.Status) bool {
	switch status {
	case
		corestatus.Available,
		corestatus.Busy,
		corestatus.Error:
		return true
	default:
		return false
	}
}

// SetStatus sets the status of the model.
//
// The following error types can be expected to be returned:
// - [modelerrors.InvalidModelStatus]: When the status to be set is not valid.
func (s *ModelService) SetStatus(ctx context.Context, params model.SetStatusArg) error {
	if !validSettableModelStatus(params.Status) {
		return internalerrors.Errorf("model status %q: %w", params.Status, modelerrors.InvalidModelStatus)
	}

	return s.modelSt.SetStatus(ctx, params)
}
