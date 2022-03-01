package guide

import (
	`github.com/goexl/gex`
)

var _ = pweEnable

func pweEnable() (int, error) {
	return gex.Exec(`redis`, gex.Pwe())
}
