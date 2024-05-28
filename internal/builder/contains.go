package builder

import (
	"github.com/goexl/gex/internal/builder/internal"
	"github.com/goexl/gex/internal/internal/checker"
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
)

type Contains struct {
	*internal.Check

	command *Command
	check   *Check
	params  *param.Check
	target  string
}

func NewContains(command *Command, check *Check, target string) (contains *Contains) {
	contains = new(Contains)
	contains.params = param.NewCheck(core.OperatorAnd)
	contains.Check = internal.NewCheck(contains.params)
	contains.command = command
	contains.check = check
	contains.target = target

	return
}

func (c *Contains) Build() (check *Check) {
	logic := core.NewLogic(c.check.operator, checker.NewContains(c.target, c.params))
	c.command.params.Logics = append(c.command.params.Logics, logic)
	check = c.check

	return
}
