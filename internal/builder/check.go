package builder

import (
	"github.com/goexl/gex/internal/internal/core"
)

type Check struct {
	command *Command
	logics  []*core.Logic
}

func NewCheck(builder *Command) *Check {
	return &Check{
		command: builder,
		logics:  make([]*core.Logic, 0),
	}
}

func (c *Check) Contains(contains string) *Contains {
	return NewContains(c.command, c, contains)
}

func (c *Check) Equal(equal string) *Equal {
	return NewEqual(c.command, c, equal)
}

func (c *Check) Filepath(pattern string) *Filepath {
	return NewFilepath(c.command, c, pattern)
}

func (c *Check) Regexp(regexp string) *Regexp {
	return NewRegexp(c.command, c, regexp)
}

func (c *Check) And(checker core.Checker) (check *Check) {
	c.logics = append(c.logics, core.NewLogic(core.OperatorAnd, checker))
	check = c

	return
}

func (c *Check) Or(checker core.Checker) (check *Check) {
	c.logics = append(c.logics, core.NewLogic(core.OperatorOr, checker))
	check = c

	return
}

func (c *Check) Build() (command *Command) {
	c.command.params.Logics = append(c.command.params.Logics, c.logics...)
	command = c.command

	return
}
