package gex

import (
	"path/filepath"
	"strings"
)

type checkerFilepathMatch struct {
	pattern string
	params  *checkerParams
	all     strings.Builder
}

func newFilepathChecker(pattern string, params *checkerParams) *checkerFilepathMatch {
	return &checkerFilepathMatch{
		pattern:  pattern,
		params: params,
	}
}

func (pmc *checkerFilepathMatch) Check(line string) (checked bool, err error) {
	if pmc.params.cache {
		pmc.all.WriteString(line)
	}

	if checked, err = filepath.Match(pmc.pattern, line); nil != err {
		return
	}
	if !checked && pmc.params.cache {
		checked, err = filepath.Match(pmc.pattern, pmc.all.String())
	}

	return
}
