package guide

import (
	"github.com/goexl/gex"
)

var _ = env

func env() (int, error) {
	return gex.New("redis").Env(`USERNAME`, `storezhang`).Build().Exec()
}
