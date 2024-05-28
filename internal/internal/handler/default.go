package handler

import (
	"os"

	"github.com/goexl/gex/internal/internal/core"
)

var _ core.Handler = (*Default)(nil)

type Default struct {
	process *os.Process
	code    int
}

func NewDefault() *Default {
	return &Default{
		process: nil,
		code:    0,
	}
}

func (d Default) Pid() int {
	return d.code
}

func (d Default) Kill() error {
	return d.process.Kill()
}

func (d Default) Code() int {
	return d.code
}
