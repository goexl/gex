package core

import (
	"github.com/goexl/gex/internal/core"
)

type Check struct {
	Operator core.Operator
	Checker  core.Checker
}

func NewCheck(operator core.Operator, checker core.Checker) *Check {
	return &Check{
		Operator: operator,
		Checker:  checker,
	}
}
