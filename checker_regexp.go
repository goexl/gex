package exec

import (
	`regexp`
	`strings`
)

type regexpChecker struct {
	regexp string

	options *checkerOptions
	all     strings.Builder
}

func (r *regexpChecker) Check(line string) (checked bool, err error) {
	if r.options.cache {
		r.all.WriteString(line)
	}

	if checked, err = regexp.MatchString(r.regexp, line); nil != err {
		return
	}
	if !checked && r.options.cache {
		checked, err = regexp.MatchString(r.regexp, r.all.String())
	}

	return
}
