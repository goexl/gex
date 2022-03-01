package guide

import (
	`github.com/goexl/gex`
)

var _ = sync

func sync() (int, error) {
	return gex.Exec(`redis`, gex.Sync())
}
