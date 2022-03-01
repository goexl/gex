package guide

import (
	`github.com/goexl/gex`
)

var _ = quiet

func quiet() (int, error) {
	return gex.Exec(`redis`, gex.Quiet())
}
