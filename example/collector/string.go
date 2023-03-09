package collector

import (
	"github.com/goexl/gex"
)

var _ = str

func str() (code int, err error) {
	output := ``
	code, err = gex.Exec(`redis`, gex.StringCollector(&output))

	return
}
