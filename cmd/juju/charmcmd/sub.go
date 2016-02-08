// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package charmcmd

import (
	"github.com/juju/cmd"
	"github.com/juju/errors"
)

var registeredSubCommands []func(CharmstoreSpec) cmd.Command

// RegisterSubCommand adds the provided func to the set of those that
// will be called when the juju command runs. Each returned command will
// be registered with the identified "juju" sub-supercommand.
func RegisterSubCommand(newCommand func(CharmstoreSpec) cmd.Command) {
	registeredSubCommands = append(registeredSubCommands, newCommand)
}

// CommandBase is the type that should be embedded in "juju charm"
// sub-commands.
type CommandBase interface {
	cmd.Command

	// Connect connects to the charm store and returns a client.
	Connect() (CharmstoreClient, error)
}

// NewCommandBase returns a new CommandBase.
func NewCommandBase(spec CharmstoreSpec) CommandBase {
	return &commandBase{
		spec: newCharmstoreSpec(),
	}
}

type commandBase struct {
	cmd.Command
	spec CharmstoreSpec
}

// Connect implements CommandBase.
func (c *commandBase) Connect() (CharmstoreClient, error) {
	if c.spec == nil {
		return nil, errors.Errorf("missing charm store spec")
	}
	client, err := c.spec.Connect()
	if err != nil {
		return nil, errors.Trace(err)
	}

	return client, nil
}
