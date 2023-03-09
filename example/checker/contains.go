package checker

import (
	"github.com/goexl/gex"
)

var _ = contains

func contains() (int, error) {
	return gex.New("github").Checker().Contains("FastGithub启动完成").Build().Exec()
}
