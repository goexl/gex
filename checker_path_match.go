package gex

import (
	`path/filepath`
	`strings`
)

type pathMatchChecker struct {
	pattern string

	options *checkerOptions
	all     strings.Builder
}

func (pmc *pathMatchChecker) Check(line string) (checked bool, err error) {
	if pmc.options.cache {
		pmc.all.WriteString(line)
	}

	if checked, err = filepath.Match(pmc.pattern, line); nil != err {
		return
	}
	if !checked && pmc.options.cache {
		checked, err = filepath.Match(pmc.pattern, pmc.all.String())
	}

	return
}
