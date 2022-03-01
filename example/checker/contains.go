package checker

import (
	`github.com/goexl/gex`
)

var _ = contains

func contains() (int, error) {
	return gex.Exec(`github`, gex.ContainsChecker(`FastGithub启动完成`))
}
