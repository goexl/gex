package guide

import (
	"github.com/goexl/gex"
)

var _ = pweEnable

func pweEnable() (int, error) {
	return gex.New("redis").Pwe().Build().Exec()
}
