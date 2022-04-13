package pipe

import (
	"github.com/goexl/gex"
)

var _ = envsParse

func envsParse() (int, error) {
	return gex.Exec(`redis`, gex.Pipe(
		`echo`,
		gex.Args(`www.163.com`),
		gex.StringEnvs(`USERNAME=storezhang`),
	))
}
