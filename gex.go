package gex

import (
	"github.com/goexl/gex/internal/builder"
)

// New 合建命令
func New(command string) *builder.Command {
	return builder.NewCommand(command)
}
