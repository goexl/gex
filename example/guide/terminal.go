package guide

import (
	`github.com/storezhang/gex`
)

var _ = terminal

func terminal() (int, error) {
	return gex.Exec(`redis`, gex.Terminal())
}
