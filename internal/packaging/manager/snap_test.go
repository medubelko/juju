// Copyright 2019 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package manager_test

import (
	"os"
	"os/exec"
	"strings"

	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/internal/packaging/commands"
	"github.com/juju/juju/internal/packaging/manager"
)

var _ = gc.Suite(&SnapSuite{})

type SnapSuite struct {
	testing.IsolationSuite
}

func (s *SnapSuite) TestInstall(c *gc.C) {
	const expected = `juju 2.6.6 from Canonical✓ installed`

	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte(expected), nil)

	paccmder := commands.NewSnapPackageCommander()
	pacman := manager.NewSnapPackageManager()
	err := pacman.Install("juju")
	c.Assert(err, gc.IsNil)

	cmd := <-cmdChan
	c.Assert(cmd.Args, gc.DeepEquals, strings.Fields(paccmder.InstallCmd("juju")))
}

func (s *SnapSuite) TestInstallWithMountFailure(c *gc.C) {
	const minRetries = 3
	var calls int
	state := os.ProcessState{}
	cmdError := &exec.ExitError{ProcessState: &state}
	s.PatchValue(&manager.SnapAttempts, minRetries)
	s.PatchValue(&manager.Delay, testing.ShortWait)
	s.PatchValue(&manager.ProcessStateSys, func(*os.ProcessState) interface{} {
		return mockExitStatuser(1) // retry each time.
	})
	s.PatchValue(&manager.CommandOutput, func(cmd *exec.Cmd) ([]byte, error) {
		calls++
		output := `
cannot perform the following tasks:
 - Mount snap "juju-db" (29) ([start snap-jujux2ddb-29.mount] failed with exit status 1: Job failed. See "journalctl -xe" for details.
)`[1:]
		return []byte(output), cmdError
	})

	pacman := manager.NewSnapPackageManager()
	err := pacman.Install("juju")
	c.Assert(err, gc.ErrorMatches, `packaging command failed: attempt count exceeded: .*`)
	c.Assert(calls, gc.Equals, minRetries)
}

func (s *SnapSuite) TestInstallWithUDevFailure(c *gc.C) {
	const minRetries = 3
	var calls int
	state := os.ProcessState{}
	cmdError := &exec.ExitError{ProcessState: &state}
	s.PatchValue(&manager.SnapAttempts, minRetries)
	s.PatchValue(&manager.Delay, testing.ShortWait)
	s.PatchValue(&manager.ProcessStateSys, func(*os.ProcessState) interface{} {
		return mockExitStatuser(2) // retry each time.
	})
	s.PatchValue(&manager.CommandOutput, func(cmd *exec.Cmd) ([]byte, error) {
		calls++
		output := `
error: cannot perform the following tasks:
 - Setup snap "snapd" (12159) security profiles (cannot reload udev rules: exit status 2
udev output:
)`[1:]
		return []byte(output), cmdError
	})

	pacman := manager.NewSnapPackageManager()
	err := pacman.Install("juju")
	c.Assert(err, gc.ErrorMatches, `packaging command failed: attempt count exceeded: .*`)
	c.Assert(calls, gc.Equals, minRetries)
}

func (s *SnapSuite) TestInstallWithFailureAndNonMatchingOutput(c *gc.C) {
	const minRetries = 3
	var calls int
	state := os.ProcessState{}
	cmdError := &exec.ExitError{ProcessState: &state}
	s.PatchValue(&manager.SnapAttempts, minRetries)
	s.PatchValue(&manager.Delay, testing.ShortWait)
	s.PatchValue(&manager.ProcessStateSys, func(*os.ProcessState) interface{} {
		return mockExitStatuser(1) // retry each time.
	})
	s.PatchValue(&manager.CommandOutput, func(cmd *exec.Cmd) ([]byte, error) {
		calls++
		// Replace the command path and args so it's a no-op.
		cmd.Path = ""
		cmd.Args = []string{"version"}
		// Call the real cmd.CombinedOutput to simulate better what
		// happens in production. See also http://pad.lv/1394524.
		output, _ := cmd.CombinedOutput()
		return output, cmdError
	})

	pacman := manager.NewSnapPackageManager()
	err := pacman.Install("juju")
	c.Assert(err, gc.ErrorMatches, `packaging command failed: exit status .*`)
	c.Assert(calls, gc.Equals, 1)
}

func (s *SnapSuite) TestInstallWithoutFailure(c *gc.C) {
	const minRetries = 3
	var calls int
	state := os.ProcessState{}
	cmdError := &exec.ExitError{ProcessState: &state}
	s.PatchValue(&manager.SnapAttempts, minRetries)
	s.PatchValue(&manager.Delay, testing.ShortWait)
	s.PatchValue(&manager.ProcessStateSys, func(*os.ProcessState) interface{} {
		return mockExitStatuser(0) // retry each time.
	})
	s.PatchValue(&manager.CommandOutput, func(cmd *exec.Cmd) ([]byte, error) {
		calls++
		// Replace the command path and args so it's a no-op.
		cmd.Path = ""
		cmd.Args = []string{"version"}
		// Call the real cmd.CombinedOutput to simulate better what
		// happens in production. See also http://pad.lv/1394524.
		output, _ := cmd.CombinedOutput()
		return output, cmdError
	})

	pacman := manager.NewSnapPackageManager()
	_ = pacman.Install("juju")
	c.Assert(calls, gc.Equals, 1)
}

func (s *SnapSuite) TestInstallWithDNSFailure(c *gc.C) {
	const minRetries = 3
	var calls int
	state := os.ProcessState{}
	cmdError := &exec.ExitError{ProcessState: &state}
	s.PatchValue(&manager.SnapAttempts, minRetries)
	s.PatchValue(&manager.Delay, testing.ShortWait)
	s.PatchValue(&manager.ProcessStateSys, func(*os.ProcessState) interface{} {
		return mockExitStatuser(100) // retry each time.
	})
	s.PatchValue(&manager.CommandOutput, func(cmd *exec.Cmd) ([]byte, error) {
		calls++
		// Replace the command path and args so it's a no-op.
		cmd.Path = ""
		cmd.Args = []string{"version"}
		// Call the real cmd.CombinedOutput to simulate better what
		// happens in production. See also http://pad.lv/1394524.
		output, _ := cmd.CombinedOutput()
		return output, cmdError
	})

	pacman := manager.NewSnapPackageManager()
	_ = pacman.Install("juju")
	c.Assert(calls, gc.Equals, 1)
}

func (s *SnapSuite) TestInstallForUnknownPackage(c *gc.C) {
	const minRetries = 3
	s.PatchValue(&manager.SnapAttempts, minRetries)
	s.PatchValue(&manager.Delay, testing.ShortWait)

	const expected = `error: snap "foo" not found`

	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte(expected), s.mockExitError(1))

	paccmder := commands.NewSnapPackageCommander()
	pacman := manager.NewSnapPackageManager()
	err := pacman.Install("foo")
	c.Assert(err, gc.ErrorMatches, ".*unable to locate package")

	cmd := <-cmdChan
	c.Assert(cmd.Args, gc.DeepEquals, strings.Fields(paccmder.InstallCmd("foo")))
}

func (s *SnapSuite) TestInstalledChannel(c *gc.C) {
	const expected = `name:      juju
summary:   juju client
publisher: Canonical✓
contact:   https://jaas.ai/
license:   unset
description: |
  Juju is an open source modelling tool for operating software in the cloud.  Juju allows you to
  ...

  https://discourse.jujucharms.com/
  https://docs.jujucharms.com/
  https://github.com/juju/juju
commands:
  - juju
snap-id:      e2CPHpB1fUxcKtCyJTsm5t3hN9axJ0yj
tracking:     2.8/bleeding-edge
refresh-date: today at 15:58 BST
channels:
  stable:        2.6.6                     2019-07-31 (8594) 68MB classic
  candidate:     ↑
  beta:          ↑
  edge:          2.7-beta1+develop-93d21f2 2019-08-19 (8756) 75MB classic
  2.6/stable:    2.6.6                     2019-07-31 (8594) 68MB classic
  ...
  2.3/beta:      ↑
  2.3/edge:      2.3.10+2.3-41313d1        2019-03-25 (7080) 55MB classic
installed:       2.6.6                                (8594) 68MB classic
`
	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte(expected), nil)

	pacman := manager.NewSnapPackageManager()
	channel := pacman.InstalledChannel("juju")
	c.Assert(channel, gc.Equals, "2.8/bleeding-edge")

	setCmd := <-cmdChan
	c.Assert(setCmd.Args, gc.DeepEquals, []string{"snap", "info", "juju"})
}

func (s *SnapSuite) TestInstalledChannelForNotInstalledSnap(c *gc.C) {
	const expected = `name:      juju
summary:   juju client
publisher: Canonical✓
contact:   https://jaas.ai/
license:   unset
description: |
  Juju is an open source modelling tool for operating software in the cloud.  Juju allows you to
  ...

  https://discourse.jujucharms.com/
  https://docs.jujucharms.com/
  https://github.com/juju/juju
commands:
  - juju
snap-id:      e2CPHpB1fUxcKtCyJTsm5t3hN9axJ0yj
refresh-date: today at 15:58 BST
channels:
  stable:        2.6.6                     2019-07-31 (8594) 68MB classic
  candidate:     ↑
  beta:          ↑
  edge:          2.7-beta1+develop-93d21f2 2019-08-19 (8756) 75MB classic
  2.6/stable:    2.6.6                     2019-07-31 (8594) 68MB classic
  ...
  2.3/beta:      ↑
  2.3/edge:      2.3.10+2.3-41313d1        2019-03-25 (7080) 55MB classic
installed:       2.6.6                                (8594) 68MB classic
`
	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte(expected), nil)

	pacman := manager.NewSnapPackageManager()
	channel := pacman.InstalledChannel("juju")
	c.Assert(channel, gc.Equals, "")

	setCmd := <-cmdChan
	c.Assert(setCmd.Args, gc.DeepEquals, []string{"snap", "info", "juju"})
}

func (s *SnapSuite) TestChangeChannel(c *gc.C) {
	const expected = `lxd (candidate) 4.0.0 from Canonical✓ refreshed`
	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte(expected), nil)

	pacman := manager.NewSnapPackageManager()
	err := pacman.ChangeChannel("lxd", "latest/candidate")
	c.Assert(err, jc.ErrorIsNil)

	setCmd := <-cmdChan
	c.Assert(setCmd.Args, gc.DeepEquals, []string{"snap", "refresh", "--channel", "latest/candidate", "lxd"})
}

func (s *SnapSuite) TestChangeChannelForNotInstalledSnap(c *gc.C) {
	const expected = `snap "lxd" is not installed`
	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte(expected), nil)

	pacman := manager.NewSnapPackageManager()
	err := pacman.ChangeChannel("lxd", "latest/candidate")
	c.Assert(err, gc.ErrorMatches, "snap not installed")

	setCmd := <-cmdChan
	c.Assert(setCmd.Args, gc.DeepEquals, []string{"snap", "refresh", "--channel", "latest/candidate", "lxd"})
}

func (s *SnapSuite) mockExitError(code int) error {
	err := &exec.ExitError{ProcessState: new(os.ProcessState)}
	s.PatchValue(&manager.ProcessStateSys, func(*os.ProcessState) interface{} {
		return mockExitStatuser(code)
	})
	return err
}
