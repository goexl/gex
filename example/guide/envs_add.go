package guide

import (
	`github.com/storezhang/gex`
)

var _ = envsAdd

func envsAdd() (int, error) {
	return gex.Exec(`redis`, gex.Envs(gex.NewEnv(`USERNAME`, `storezhang`)))
}
