package guide

import (
	`github.com/storezhang/gex`
)

var _ = pweDisable

func pweDisable() (int, error) {
	return gex.Exec(`redis`, gex.DisablePwe())
}
