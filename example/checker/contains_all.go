package checker

import (
	"github.com/goexl/gex"
)

var _ = containsAll

func containsAll() (int, error) {
	return gex.Exec(`github`, gex.ContainsAllChecker([]string{`FastGithub启动完成`, `另外一个检查`}))
}
