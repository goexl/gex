package param

import (
	"github.com/goexl/gox"
)

type Line struct {
	Stream string
	Total  *int

	_ gox.CannotCopy
}

func NewLine(total *int) *Line {
	return &Line{
		Total: total,
	}
}
