package param

import (
	"github.com/goexl/gox"
)

type Environment struct {
	Environments []string
	System       bool

	_ gox.CannotCopy
}

func NewEnvironment() *Environment {
	return &Environment{
		Environments: make([]string, 0, 1),
		System:       true,
	}
}
