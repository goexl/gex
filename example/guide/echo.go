package guide

import (
	"github.com/goexl/gex"
)

var _ = echo

func echo() (int, error) {
	return gex.New("redis").Echo().Build().Exec()
}
