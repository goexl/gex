package checker

import (
	"regexp"
	"strings"

	"github.com/goexl/gex/internal/param"
)

type Regexp struct {
	regexp string
	params *param.Check
	all    strings.Builder
}

func NewRegexp(regexp string, params *param.Check) *Regexp {
	return &Regexp{
		regexp: regexp,
		params: params,
	}
}

func (r *Regexp) Check(line string) (checked bool, err error) {
	if r.params.Cache {
		r.all.WriteString(line)
	}

	if checked, err = regexp.MatchString(r.regexp, line); nil != err {
		return
	}
	if !checked && r.params.Cache {
		checked, err = regexp.MatchString(r.regexp, r.all.String())
	}

	return
}
