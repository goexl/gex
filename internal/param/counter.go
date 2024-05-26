package param

import (
	"github.com/goexl/gox"
)

type Counter struct {
	gox.CannotCopy

	Stream string
}

func NewCounter() *Counter {
	return new(Counter)
}
