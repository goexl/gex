package param

import (
	"github.com/goexl/gox"
)

type Checker struct {
	gox.CannotCopy

	Cache bool
}

func NewChecker() *Checker {
	return &Checker{
		Cache: false,
	}
}
