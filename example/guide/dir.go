package guide

import (
	`github.com/storezhang/gex`
)

var _ = dir

func dir() (int, error) {
	return gex.Exec(`redis`, gex.Dir(`/opt/`))
}
