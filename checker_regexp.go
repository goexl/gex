package gex

import (
	`regexp`
	`strings`
)

type regexpChecker struct {
	regexp string

	options *checkerOptions
	all     strings.Builder
}

func (rc *regexpChecker) Check(line string) (checked bool, err error) {
	if rc.options.cache {
		rc.all.WriteString(line)
	}

	if checked, err = regexp.MatchString(rc.regexp, line); nil != err {
		return
	}
	if !checked && rc.options.cache {
		checked, err = regexp.MatchString(rc.regexp, rc.all.String())
	}

	return
}
