package builder

import (
	"fmt"

	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/internal/core"
)

type Collect struct {
	command    *Command
	collectors []core.Collector
}

func NewCollect(command *Command) *Collect {
	return &Collect{
		command:    command,
		collectors: make([]core.Collector, 0),
	}
}

func (c *Collect) Strings(target *[]string) *Strings {
	return NewStrings(c.command, c, target)
}

func (c *Collect) String(target *string) *String {
	return NewString(c.command, c, target)
}

func (c *Collect) By(collector core.Collector) (collect *Collect) {
	c.command.params.Collectors[fmt.Sprintf(constant.Point, collector)] = collector
	collect = c

	return
}

func (c *Collect) Build() (command *Command) {
	for _, collector := range c.collectors {
		c.command.params.Collectors[fmt.Sprintf(constant.Point, collector)] = collector
	}
	c.command.params.Async = true
	command = c.command

	return
}
