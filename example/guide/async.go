package guide

import (
	"github.com/goexl/gex"
)

var _ = async

func async() (int, error) {
	return gex.New("redis").Async().Build().Exec()
}
