// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package resources

import (
	"testing"

	gc "gopkg.in/check.v1"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package resources -destination resource_opener_mock_test.go github.com/juju/juju/core/resource Opener
//go:generate go run go.uber.org/mock/mockgen -typed -package resources -destination service_mock_test.go github.com/juju/juju/apiserver/internal/handlers/resources ResourceServiceGetter,ApplicationServiceGetter,ApplicationService,ResourceService,ResourceOpenerGetter

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}
