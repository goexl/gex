package collector

import (
	"github.com/goexl/gex"
)

var _ = str

func str() (code int, err error) {
	output := ""
	code, err = gex.New("redis").Collector().String(&output).Build().Exec()

	return
}
