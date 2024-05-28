package internal

import (
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
)

type Check struct {
	params *param.Check
}

func NewCheck(params *param.Check) *Check {
	return &Check{
		params: params,
	}
}

func (c *Check) And() (check *Check) {
	c.params.Operator = core.OperatorAnd
	check = c

	return
}

func (c *Check) Or() (check *Check) {
	c.params.Operator = core.OperatorOr
	check = c

	return
}

func (c *Check) Cached() (check *Check) {
	c.params.Cache = true
	check = c

	return
}
