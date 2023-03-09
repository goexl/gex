package checker

import (
	"github.com/goexl/gex"
)

var _ = equal

func equal() (int, error) {
	return gex.Exec(`github`, gex.EqualChecker(`FastGithub启动完成`))
}
