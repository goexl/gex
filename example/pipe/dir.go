package pipe

import (
	`github.com/storezhang/gex`
)

var _ = dir

func dir() (int, error) {
	return gex.Exec(`ping`, gex.Pipe(`echo`, gex.Args(`www.163.com`), gex.Dir(`/opt/`)))
}
