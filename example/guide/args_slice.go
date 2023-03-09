package guide

import (
	"github.com/goexl/gex"
)

var _ = slice

func slice() (int, error) {
	return gex.Exec(`ping`, gex.Args(`www.163.com`, `-c`, 10))
}
