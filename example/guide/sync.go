package guide

import (
	`github.com/storezhang/gex`
)

var _ = sync

func sync() (int, error) {
	return gex.Exec(`redis`, gex.Sync())
}
