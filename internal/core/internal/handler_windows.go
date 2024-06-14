//go:build windows

package internal

import (
	"os"
)

func (h *Handler) kill() error {
	return h.cmd.Process.Signal(os.Interrupt)
}
