package checker

import (
	"github.com/goexl/gex"
)

var _ = containsAny

func containsAny() (int, error) {
	return gex.Exec(`github`, gex.ContainsAnyChecker([]string{`FastGithub启动完成`, `另外一个检查`}))
}
