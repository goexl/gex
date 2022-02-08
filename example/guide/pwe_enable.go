package guide

import (
	`github.com/storezhang/gex`
)

var _ = pweEnable

func pweEnable() (int, error) {
	return gex.Exec(`redis`, gex.Pwe())
}
