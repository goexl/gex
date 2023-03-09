package checker

import (
	"github.com/goexl/gex"
)

var _ = equal

func equal() (int, error) {
	return gex.New("github").Checker().Equal("FastGithub启动完成").Build().Exec()
}
