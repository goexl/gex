package internal

import (
	"github.com/goexl/gex/internal/param"
	"github.com/goexl/gox"
)

type Trim struct {
	collect *Collect
	params  *param.Trim

	_ gox.CannotCopy
}

func NewTrim(collect *Collect) (trim *Trim) {
	return &Trim{
		collect: collect,
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

func (t *Trim) Build() (collector *Collect) {
	t.collect.params.Trim = t.params
	collector = t.collect

	return
}
