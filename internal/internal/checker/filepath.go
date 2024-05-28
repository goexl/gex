package checker

import (
	"path/filepath"
	"strings"

	"github.com/goexl/gex/internal/param"
)

type Filepath struct {
	pattern string
	params  *param.Check
	all     strings.Builder
}

func NewFilepath(pattern string, params *param.Check) *Filepath {
	return &Filepath{
		pattern: pattern,
		params:  params,
	}
}

func (f *Filepath) Check(line string) (checked bool, err error) {
	if f.params.Cache {
		f.all.WriteString(line)
	}

	if checked, err = filepath.Match(f.pattern, line); nil != err {
		return
	}
	if !checked && f.params.Cache {
		checked, err = filepath.Match(f.pattern, f.all.String())
	}

	return
}
