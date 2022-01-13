package gex

import (
	`path/filepath`
)

type pathMatchChecker struct {
	pattern string
}

func (pmc *pathMatchChecker) check(all string, line string) (checked bool, err error) {
	if checked, err = filepath.Match(pmc.pattern, line); nil != err {
		return
	}

	if !checked && `` != all {
		checked, err = filepath.Match(pmc.pattern, all)
	}

	return
}
