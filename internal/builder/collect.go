package builder

import (
	"fmt"

	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/internal/core"
)

type Collect struct {
	command *Command
}

func NewCollect(command *Command) *Collect {
	return &Collect{
		command: command,
	}
}

func (c *Collect) Strings(target *[]string) *Strings {
	return NewStrings(c.command, c, target)
}

func (c *Collect) String(target *string) *String {
	return NewString(c.command, c, target)
}

func (c *Collect) By(collector core.Collector) (command *Command) {
	c.command.params.Collectors[fmt.Sprintf(constant.Point, collector)] = collector
	command = c.command

	return
}
