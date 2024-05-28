package param

import (
	"github.com/goexl/gox"
)

type Collect struct {
	Stream string
	Max    int
	Trim   *Trim

	_ gox.CannotCopy
}

func NewCollect() *Collect {
	return &Collect{
		Max:  1024,
		Trim: NewTrim(),
	}
}

func (c *Collect) Process(line string) (final string) {
	final = line
	if nil != c.Trim {
		final = c.Trim.Process(line)
	}

	return
}
