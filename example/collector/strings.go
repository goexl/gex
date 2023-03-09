package collector

import (
	"github.com/goexl/gex"
)

var _ = strings

func strings() (code int, err error) {
	output := make([]string, 0, 1)
	code, err = gex.New("redis").Collector().Strings(&output).Build().Exec()

	return
}
