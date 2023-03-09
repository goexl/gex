package guide

import (
	"os"

	"github.com/goexl/gex"
)

var _ = stdoutWriter

func stdoutWriter() (code int, err error) {
	var output *os.File
	if output, err = os.Open(`/var/log/redis.log`); nil != err {
		return
	}
	code, err = gex.Exec(`redis`, gex.Stdout(output))

	return
}
