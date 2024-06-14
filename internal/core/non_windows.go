//go:build !windows
// +build !windows

package core

import (
	"syscall"
)

func (c *Command) setup() {
	c.cmd.SysProcAttr = new(syscall.SysProcAttr)
	c.cmd.SysProcAttr.Setpgid = true
}
