// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package common

import (
	"github.com/juju/errors"
)

// Conf is responsible for defining services. Its fields
// represent elements of a service configuration.
type Conf struct {
	// Desc is the init service's description.
	Desc string

	// Env holds the environment variables that will be set when the
	// command runs.
	// Currently not used on Windows
	Env map[string]string

	// Limit holds the ulimit values that will be set when the command runs.
	// Currently not used on Windows
	Limit map[string]string

	// TODO(ericsnow) Cmd -> ExecStart.

	// Cmd is the command (with arguments) that will be run.
	// The command will be restarted if it exits with a non-zero exit code.
	Cmd string

	// Out, if set, will redirect output to that path.
	Out string

	// TODO(ericsnow) Eliminate InitDir.

	// InitDir is the folder in which the init script should be written
	// defaults to "/etc/init" on Ubuntu
	// Currently not used on Windows
	InitDir string

	// TODO(ericsnow) Turn ExtraScript into ExecStartPre.

	// ExtraScript allows to insert script before command execution
	ExtraScript string
}

// Validate checks the conf's values for correctness.
func (c Conf) Validate() error {
	if c.Desc == "" {
		return errors.New("missing Desc")
	}

	if c.Cmd == "" {
		return errors.New("missing Cmd")
	}

	return nil
}
