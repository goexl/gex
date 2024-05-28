package param

import (
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gox"
)

type Check struct {
	Cache    bool
	Operator core.Operator

	_ gox.CannotCopy
}

func NewCheck(operator core.Operator) *Check {
	return &Check{
		Cache:    false,
		Operator: operator,
	}
}
