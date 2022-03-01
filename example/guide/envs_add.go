package guide

import (
	`github.com/goexl/gex`
)

var _ = envsAdd

func envsAdd() (int, error) {
	return gex.Exec(`redis`, gex.Envs(gex.NewEnv(`USERNAME`, `storezhang`)))
}
