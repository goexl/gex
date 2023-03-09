package checker

import (
	"github.com/goexl/gex"
)

var _ = filepath

func filepath() (int, error) {
	return gex.New("github").Checker().Filepath("*FastGithub*").Build().Exec()
}
