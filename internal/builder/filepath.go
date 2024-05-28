package builder

import (
	"github.com/goexl/gex/internal/builder/internal"
	"github.com/goexl/gex/internal/internal/checker"
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
)

type Filepath struct {
	*internal.Check

	command *Command
	check   *Check
	params  *param.Check
	target  string
}

func NewFilepath(command *Command, check *Check, target string) (filepath *Filepath) {
	filepath = new(Filepath)
	filepath.params = param.NewCheck(core.OperatorAnd)
	filepath.Check = internal.NewCheck(filepath.params)
	filepath.command = command
	filepath.check = check
	filepath.target = target

	return
}

func (c *Filepath) Build() (check *Check) {
	logic := core.NewLogic(c.check.operator, checker.NewFilepath(c.target, c.params))
	c.command.params.Logics = append(c.command.params.Logics, logic)
	check = c.check

	return
}
