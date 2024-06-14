//go:build linux

package internal

import (
	"syscall"
)

func (h *Handler) kill() error {
	return syscall.Kill(-h.cmd.Process.Pid, syscall.SIGKILL)
}