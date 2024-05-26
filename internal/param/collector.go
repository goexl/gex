package param

import (
	"github.com/goexl/gox"
)

type Collector struct {
	gox.CannotCopy

	Stream string
	Max    int
	Trim   *Trim
}

func NewCollector() *Collector {
	return &Collector{
		Max:  1024,
		Trim: NewTrim(),
	}
}

func (c *Collector) Process(line string) (final string) {
	final = line
	if nil != c.Trim {
		final = c.Trim.process(line)
	}

	return
}
