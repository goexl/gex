//go:build linux

package core

import (
	"syscall"
)

func (c *Command) setup() {
	c.cmd.SysProcAttr = new(syscall.SysProcAttr)
	c.cmd.SysProcAttr.Setpgid = true
}
