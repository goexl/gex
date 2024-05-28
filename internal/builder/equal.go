package builder

import (
	"github.com/goexl/gex/internal/builder/internal"
	"github.com/goexl/gex/internal/internal/checker"
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
)

type Equal struct {
	*internal.Check

	command *Command
	check   *Check
	params  *param.Check
	target  string
}

func NewEqual(command *Command, check *Check, target string) (equal *Equal) {
	equal = new(Equal)
	equal.params = param.NewCheck(core.OperatorAnd)
	equal.Check = internal.NewCheck(equal.params)
	equal.command = command
	equal.check = check
	equal.target = target

	return
}

func (c *Equal) Build() (check *Check) {
	logic := core.NewLogic(c.params.Operator, checker.NewEqual(c.target, c.params))
	c.command.params.Logics = append(c.command.params.Logics, logic)
	check = c.check

	return
}
