package param

import (
	"strings"

	"github.com/goexl/gox"
)

type Trim struct {
	Left  string
	Right string
	All   string
	Space bool

	_ gox.CannotCopy
}

func NewTrim() *Trim {
	return &Trim{
		Left:  "",
		Right: "",
		Space: false,
	}
}

func (t *Trim) Process(line string) (final string) {
	final = line
	if "" != t.Left {
		final = strings.TrimLeft(final, t.Left)
	}
	if "" != t.Right {
		final = strings.TrimRight(final, t.Right)
	}
	if t.Space {
		final = strings.TrimSpace(final)
	}

	return
}
