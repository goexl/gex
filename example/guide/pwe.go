package guide

import (
	"github.com/goexl/gex"
)

var _ = pwe

func pwe() (int, error) {
	return gex.New("redis").Pwe().Build().Exec()
}
