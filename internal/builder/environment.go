package builder

import (
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox"
)

type Environment struct {
	builder *Command
	params  *param.Environment
}

func NewEnvironment(builder *Command) *Environment {
	return &Environment{
		builder: builder,
		params:  param.NewEnvironment(),
	}
}

func (e *Environment) Only() *Environment {
	e.params.System = false

	return e
}

func (e *Environment) Kv(key string, value any) *Environment {
	e.params.Environments = append(e.params.Environments, gox.StringBuilder(key, constant.Equal, value).String())

	return e
}

func (e *Environment) String(environments ...string) *Environment {
	for _, environment := range environments {
		e.params.Environments = append(e.params.Environments, environment)
	}

	return e
}

func (e *Environment) Build() *Command {
	e.builder.params.Environment = e.params

	return e.builder
}
