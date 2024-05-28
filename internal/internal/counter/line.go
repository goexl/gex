package counter

import (
	"github.com/goexl/gex/internal/param"
)

type Line struct {
	params *param.Line
}

func NewLine(params *param.Line) *Line {
	return &Line{
		params: params,
	}
}

func (l *Line) Count(line string, stream string) (err error) {
	if "" != line && ("" == l.params.Stream || stream == l.params.Stream) {
		*l.params.Total++
	}

	return
}
