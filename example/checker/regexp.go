package guide

import (
	`github.com/storezhang/gex`
)

var _ = regexp

func regexp() (int, error) {
	return gex.Exec(`github`, gex.RegexpChecker(`FastGithub*`))
}
