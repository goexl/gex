package param

import (
	"github.com/goexl/gex/internal/internal/core"
)

type Check struct {
	Cache    bool
	Operator core.Operator
}

func NewCheck(operator core.Operator) *Check {
	return &Check{
		Cache:    false,
		Operator: operator,
	}
}
