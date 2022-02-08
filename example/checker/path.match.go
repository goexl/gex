package guide

import (
	`github.com/storezhang/gex`
)

var _ = pathMatch

func pathMatch() (int, error) {
	return gex.Exec(`github`, gex.PathMatchChecker(`*FastGithub*`))
}
