package pipe

import (
	`github.com/storezhang/gex`
)

var _ = usage

func usage() (int, error) {
	return gex.Exec(`ping`, gex.Pipe(`echo`, gex.Args(`www.163.com`)))
}
