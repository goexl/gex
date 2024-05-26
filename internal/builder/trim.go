package builder

import (
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox"
)

type Trim struct {
	gox.CannotCopy

	builder *Collector
	params  *param.Trim
}

func NewTrim(builder *Collector) (trim *Trim) {
	return &Trim{
		builder: builder,
		params:  param.NewTrim(),
	}
}

func (t *Trim) TrimSpace() (trim *Trim) {
	t.params.Space = true
	trim = t

	return
}

func (t *Trim) Trim(all string) (trim *Trim) {
	t.params.All = all
	trim = t

	return
}

func (t *Trim) Left(left string) (trim *Trim) {
	t.params.Left = left
	trim = t

	return
}

func (t *Trim) Right(right string) (trim *Trim) {
	t.params.Right = right
	trim = t

	return
}

func (t *Trim) Build() (collector *Collector) {
	t.builder.params.Trim = t.params
	collector = t.builder

	return
}
