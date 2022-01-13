package gex

import (
	`regexp`
)

type regexpChecker struct {
	regexp string
}

func (rc *regexpChecker) check(all string, line string) (checked bool, err error) {
	if checked, err = regexp.MatchString(rc.regexp, line); nil != err {
		return
	}

	if !checked && `` != all {
		checked, err = regexp.MatchString(rc.regexp, all)
	}

	return
}
