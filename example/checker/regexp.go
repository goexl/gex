package checker

import (
	"github.com/goexl/gex"
)

var _ = regexp

func regexp() (int, error) {
	return gex.New("github").Checker().Regexp("FastGithub*").Build().Exec()
}
