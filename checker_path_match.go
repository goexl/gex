package gex

import (
	`path/filepath`
	`strings`
)

type pathMatchChecker struct {
	pattern string

	cache bool
	all   strings.Builder
}

func (pmc *pathMatchChecker) check(line string) (checked bool, err error) {
	if pmc.cache {
		pmc.all.WriteString(line)
	}

	if checked, err = filepath.Match(pmc.pattern, line); nil != err {
		return
	}
	if !checked && pmc.cache {
		checked, err = filepath.Match(pmc.pattern, pmc.all.String())
	}

	return
}
