package guide

import (
	`github.com/storezhang/gex`
)

var _ = async

func async() (int, error) {
	return gex.Exec(`redis`, gex.Async())
}
