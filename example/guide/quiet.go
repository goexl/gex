package guide

import (
	`github.com/storezhang/gex`
)

var _ = quiet

func quiet() (int, error) {
	return gex.Exec(`redis`, gex.Quiet())
}
