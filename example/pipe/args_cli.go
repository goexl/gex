package pipe

import (
	`github.com/storezhang/gex`
)

var _ = cli

func cli() (int, error) {
	return gex.Exec(`ping`, gex.Pipe(`echo`, gex.Cli(`www.163.com`)))
}
