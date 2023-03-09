package pipe

import (
	"github.com/goexl/gex"
)

var _ = slice

func slice() (int, error) {
	return gex.Exec(`ping`, gex.Pipe(`echo`, gex.Args(`www.163.com`)))
}
