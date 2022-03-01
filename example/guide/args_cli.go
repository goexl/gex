package guide

import (
	`github.com/goexl/gex`
)

var _ = cli

func cli() (int, error) {
	return gex.Exec(`ping`, gex.Cli(`www.163.com -c 10`))
}
