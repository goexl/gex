package builder

import (
	"github.com/goexl/gex/internal/builder/internal"
	"github.com/goexl/gex/internal/internal/checker"
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
)

type Regexp struct {
	*internal.Check

	command *Command
	check   *Check
	params  *param.Check
	target  string
}

func NewRegexp(command *Command, check *Check, target string) (regexp *Regexp) {
	regexp = new(Regexp)
	regexp.params = param.NewCheck(core.OperatorAnd)
	regexp.Check = internal.NewCheck(regexp.params)
	regexp.command = command
	regexp.check = check
	regexp.target = target

	return
}

func (c *Regexp) Build() (check *Check) {
	logic := core.NewLogic(c.check.operator, checker.NewRegexp(c.target, c.params))
	c.command.params.Logics = append(c.command.params.Logics, logic)
	check = c.check

	return
}
