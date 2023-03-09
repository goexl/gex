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
		pattern: pattern,
		params:  params,
	}
}

func (cfm *checkerFilepathMatch) Check(line string) (checked bool, err error) {
	if cfm.params.cache {
		cfm.all.WriteString(line)
	}

	if checked, err = filepath.Match(cfm.pattern, line); nil != err {
		return
	}
	if !checked && cfm.params.cache {
		checked, err = filepath.Match(cfm.pattern, cfm.all.String())
	}

	return
}
