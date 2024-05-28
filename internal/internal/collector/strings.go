package collector

import (
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
)

var _ core.Collector = (*Strings)(nil)

type Strings struct {
	strings *[]string
	params  *param.Collect
}

func NewStrings(strings *[]string, params *param.Collect) *Strings {
	return &Strings{
		strings: strings,
		params:  params,
	}
}

func (s *Strings) Collect(line string, stream string) (err error) {
	if "" == s.params.Stream || stream == s.params.Stream {
		line = s.params.process(line)
		*s.strings = append(*s.strings, line)
	}

	return
}

func (s *Strings) Notify(_ int, _ error) {}
