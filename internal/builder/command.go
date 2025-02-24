package builder

import (
	"context"

	"github.com/goexl/args"
	"github.com/goexl/gex/internal/core"
	"github.com/goexl/gex/internal/param"
)

type Command struct {
	params *param.Command
}

func NewCommand(name string) *Command {
	return &Command{
		params: param.NewCommand(name),
	}
}

func (c *Command) Name(name string) (command *Command) {
	c.params.Name = name
	command = c

	return
}

func (c *Command) Arguments(arguments *args.Arguments) (command *Command) {
	c.params.Arguments = arguments
	command = c

	return
}

func (c *Command) Context(ctx context.Context) (command *Command) {
	c.params.Context = ctx
	command = c

	return
}

func (c *Command) Dir(dir string) (command *Command) {
	c.params.Directory = dir
	command = c

	return
}

func (c *Command) Environment() *Environment {
	return NewEnvironment(c)
}

func (c *Command) Async() (command *Command) {
	c.params.Async = true
	command = c

	return
}

func (c *Command) Pwe() (command *Command) {
	c.params.Pwe = true
	command = c

	return
}

func (c *Command) Charset(charset string) (command *Command) {
	c.params.Charset = charset
	command = c

	return
}

func (c *Command) Echo() (command *Command) {
	c.params.Echo = true
	command = c

	return
}

func (c *Command) Check() *Check {
	return NewCheck(c)
}

func (c *Command) Collect() *Collect {
	return NewCollect(c)
}

func (c *Command) Count() *Count {
	return NewCount(c)
}

func (c *Command) Notify() *Notify {
	return NewNotify(c)
}

func (c *Command) Build() *core.Command {
	return core.NewCommand(c.params)
}
