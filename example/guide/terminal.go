package guide

import (
	`github.com/goexl/gex`
)

var _ = terminal

func terminal() (int, error) {
	return gex.Exec(`redis`, gex.Terminal())
}
