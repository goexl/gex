package pipe

import (
	"github.com/goexl/gex"
)

var _ = envsAdd

func envsAdd() (int, error) {
	return gex.Exec(`ping`, gex.Pipe(
		`echo`,
		gex.Args(`www.163.com`),
		gex.Env(`USERNAME`, `storezhang`),
	))
}
