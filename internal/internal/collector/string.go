package collector

import (
	"strings"

	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/internal/core"
	"github.com/goexl/gex/internal/param"
)

var _ core.Collector = (*String)(nil)

type String struct {
	string *string
	params *param.Collect

	lines []string
}

func NewTerminal(output *string) *String {
	return NewString(output, &param.Collect{
		Max: 1024,
	})
}

func NewString(output *string, params *param.Collect) *String {
	return &String{
		string: output,
		params: params,

		lines: make([]string, 0, params.Max),
	}
}

func (s *String) Collect(line string, stream string) (err error) {
	if "" != s.params.Stream && stream != s.params.Stream {
		return
	}

	line = s.params.Process(line)
	s.lines = append(s.lines, line)
	current := len(s.lines)
	if current > s.params.Max {
		s.lines = s.lines[current-s.params.Max:]
	}

	return
}

func (s *String) Notify(_ int, _ error) {
	*s.string = strings.Join(s.lines, constant.Empty)
}
