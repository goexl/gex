package guide

import (
	"github.com/goexl/gex"
)

var _ = envDisableSystem

func envDisableSystem() (int, error) {
	return gex.Exec(`redis`, gex.DisableSystemEnv())
}
