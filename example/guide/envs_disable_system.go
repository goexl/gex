package guide

import (
	`github.com/storezhang/gex`
)

var _ = envsDisableSystem

func envsDisableSystem() (int, error) {
	return gex.Exec(`redis`, gex.DisableSystemEnvs())
}
