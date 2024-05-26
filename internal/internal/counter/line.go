package counter

import (
	"github.com/goexl/gex/internal/param"
)

type Line struct {
	total  *int
	params *param.Counter
}

func NewLine(total *int, params *param.Counter) *Line {
	return &Line{
		total:  total,
		params: params,
	}
}

func (l *Line) Count(line string, stream string) (err error) {
	if "" != line && ("" == l.params.Stream || stream == l.params.Stream) {
		*l.total++
	}

	return
}
