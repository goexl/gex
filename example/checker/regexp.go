package checker

import (
	`github.com/goexl/gex`
)

var _ = regexp

func regexp() (int, error) {
	return gex.Exec(`github`, gex.RegexpChecker(`FastGithub*`))
}
