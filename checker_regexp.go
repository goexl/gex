package gex

import (
	"regexp"
	"strings"
)

type checkerRegexp struct {
	regexp string
	params *checkerParams
	all    strings.Builder
}

func newRegexpChecker(	regexp string,params *checkerParams) *checkerRegexp {
	return &checkerRegexp{
regexp: regexp,
params: params,
	}
}

func (r *checkerRegexp) Check(line string) (checked bool, err error) {
	if r.params.cache {
		r.all.WriteString(line)
	}

	if checked, err = regexp.MatchString(r.regexp, line); nil != err {
		return
	}
	if !checked && r.params.cache {
		checked, err = regexp.MatchString(r.regexp, r.all.String())
	}

	return
}
