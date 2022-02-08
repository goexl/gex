package guide

import (
	`github.com/storezhang/gex`
)

var _ = args

func args() (int, error) {
	return gex.Exec(`ping`, gex.Args(`www.163.com`))
}
