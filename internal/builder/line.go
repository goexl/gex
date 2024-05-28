package builder

import (
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/internal/counter"
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox"
)

type Line struct {
	count  *Count
	params *param.Line

	_ gox.CannotCopy
}

func NewLine(count *Count, total *int) *Line {
	return &Line{
		count:  count,
		params: param.NewLine(total),
	}
}

func (l *Line) Stdout() (line *Line) {
	l.params.Stream = constant.Stdout
	line = l

	return
}

func (l *Line) Stderr() (line *Line) {
	l.params.Stream = constant.Stderr
	line = l

	return
}

func (l *Line) Build() (cnt *Count) {
	l.count.counters = append(l.count.counters, counter.NewLine(l.params))
	cnt = l.count

	return
}
