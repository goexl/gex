package param

import (
	"github.com/goexl/gox"
)

type Environment struct {
	gox.CannotCopy

	Environments []string
	System       bool
}

func NewEnvironment() *Environment {
	return &Environment{
		Environments: make([]string, 0, 1),
		System:       true,
	}
}
