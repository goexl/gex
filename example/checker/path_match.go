package checker

import (
	"github.com/goexl/gex"
)

var _ = pathMatch

func pathMatch() (int, error) {
	return gex.Exec(`github`, gex.PathMatchChecker(`*FastGithub*`))
}
