package builder

import (
	"github.com/goexl/gex/internal/core"
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox"
)

type Counter struct {
	gox.CannotCopy

	command  *Command
	params   *param.Counter
	counters []core.Counter
}

func NewCounter(command *Command) (counter *Counter) {
	return &Counter{
		command:  command,
		params:   param.NewCounter(),
		counters: make([]core.Counter, 0),
	}
}

func (c *Counter) Stdout() (counter *Counter) {
	c.params.Stream = constant.Stdout
	counter = c

	return
}

func (c *Counter) Stderr() (counter *Counter) {
	c.params.Stream = constant.Stderr
	counter = c

	return
}

func (c *Counter) Line(total *int) (counter *Counter) {
	c.counters = append(c.counters, counter.NewLine(total, c.params))
	counter = c

	return
}

func (c *Counter) By(by core.Counter) (counter *Counter) {
	c.counters = append(c.counters, by)
	counter = c

	return
}

func (c *Counter) Build() (command *Command) {
	c.command.params.Counters = append(c.command.params.Counters, c.counters...)
	command = c.command

	return
}
