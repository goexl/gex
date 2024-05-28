package builder

import (
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox"
)

type Environment struct {
	command *Command
	params  *param.Environment

	_ gox.CannotCopy
}

func NewEnvironment(command *Command) *Environment {
	return &Environment{
		command: command,
		params:  param.NewEnvironment(),
	}
}

func (e *Environment) Kv(key string, value any) (environment *Environment) {
	e.params.Environments = append(e.params.Environments, gox.StringBuilder(key, constant.Equal, value).String())
	environment = e

	return
}

func (e *Environment) String(environments ...string) (environment *Environment) {
	for _, env := range environments {
		e.params.Environments = append(e.params.Environments, env)
	}
	environment = e

	return
}

func (e *Environment) Build() (command *Command) {
	e.command.params.Environment = e.params
	command = e.command

	return
}
