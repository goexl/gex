package internal

import (
	"os/exec"
	"syscall"

	"github.com/goexl/gex/internal/internal/core"
)

var _ core.Handler = (*Handler)(nil)

type Handler struct {
	cmd      *exec.Cmd
	executor core.Executor
}

func NewHandler(cmd *exec.Cmd, executor core.Executor) *Handler {
	return &Handler{
		cmd:      cmd,
		executor: executor,
	}
}

func (h *Handler) Pid() int {
	return h.cmd.Process.Pid
}

func (h *Handler) Kill() (err error) {
	if se := h.cmd.Process.Signal(syscall.Signal(0)); nil == se {
		err = h.kill()
	}
	if nil == err {
		err = h.cmd.Process.Release()
	}

	return
}

func (h *Handler) Code() (code int) {
	if ee, ok := h.cmd.Err.(*exec.ExitError); ok {
		if status, sok := ee.Sys().(syscall.WaitStatus); sok {
			code = status.ExitStatus()
		}
	}

	return
}

func (h *Handler) Restart() (err error) {
	if ke := h.Kill(); nil != ke {
		err = ke
	} else if handler, ee := h.executor.Exec(); nil != ee {
		err = ee
	} else if def, ok := handler.(*Handler); ok {
		*h = *def
	}

	return
}
