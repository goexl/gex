package checker

import (
	`github.com/storezhang/gex`
)

var _ = contains

func contains() (int, error) {
	return gex.Exec(`github`, gex.ContainsChecker(`FastGithub启动完成`))
}
