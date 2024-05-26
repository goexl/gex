package gex

import (
	"github.com/goexl/gex/internal/builder"
)

var _ = New

// New 合建命令
func New(command string) *builder.Command {
	return builder.NewCommand(command)
}
