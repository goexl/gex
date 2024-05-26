package builder

import (
	"context"
	"fmt"

	"github.com/goexl/args"
	"github.com/goexl/gex"
	"github.com/goexl/gex/internal/core"
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox"
)

type Command struct {
	gox.CannotCopy

	params *param.Command
}

func NewCommand(name string) (command *Command) {
	return &Command{
		params: param.NewCommand(name),
	}
}

func (c *Command) Args(arguments *args.Arguments) (command *Command) {
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

func (c *Command) Checker() (builder *Checker) {
	c.params.Async = true
	builder = NewChecker(core.OperatorAnd, c)

	return
}

func (c *Command) Check(checker core.Checker) (command *Command) {
	c.params.Checks = append(c.params.Checks, gex.newCheck(core.OperatorAnd, checker))

	return c
}

func (c *Command) Collector() *Collector {
	return NewCollector(c)
}

func (c *Command) Collect(collector core.Collector) (command *Command) {
	c.params.Collectors[fmt.Sprintf(constant.Point, collector)] = collector
	command = c

	return
}

func (c *Command) Counter() *Counter {
	return NewCounter(c)
}

func (c *Command) Count(counter core.Counter) (command *Command) {
	c.params.Counters = append(c.params.Counters, counter)
	command = c

	return
}

func (c *Command) Notify(notifier core.Notifier) (command *Command) {
	c.params.Notifiers = append(c.params.Notifiers, notifier)
	command = c

	return
}

func (c *Command) Build() *core.Command {
	return core.NewCommand(c.params)
}
