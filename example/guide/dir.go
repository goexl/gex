package guide

import (
	"github.com/goexl/gex"
)

var _ = dir

func dir() (int, error) {
	return gex.New("redis").Dir("/opt/").Build().Exec()
}
