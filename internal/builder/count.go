package builder

import (
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gox"
)

type Count struct {
	command  *Command
	counters []core.Counter

	_ gox.CannotCopy
}

func NewCount(command *Command) (counter *Count) {
	return &Count{
		command:  command,
		counters: make([]core.Counter, 0),
	}
}

func (c *Count) Line(total *int) *Line {
	return NewLine(c, total)
}

func (c *Count) By(by core.Counter) (count *Count) {
	c.counters = append(c.counters, by)
	count = c

	return
}

func (c *Count) Build() (command *Command) {
	c.command.params.Counters = append(c.command.params.Counters, c.counters...)
	command = c.command

	return
}
