package checker

import (
	`github.com/goexl/gex`
)

var _ = usage

type checker struct{}

func newChecker() *checker {
	return new(checker)
}

func (c *checker) Check(line string) (checked bool, err error) {
	checked = `1` == line

	return
}

func usage() (int, error) {
	return gex.Exec(`github`, gex.Checker(newChecker()))
}
