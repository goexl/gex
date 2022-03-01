package guide

import (
	`github.com/goexl/gex`
)

var _ = envsDisableSystem

func envsDisableSystem() (int, error) {
	return gex.Exec(`redis`, gex.DisableSystemEnvs())
}
