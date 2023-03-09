package guide

import (
	"os"

	"github.com/goexl/gex"
)

var _ = stderrWriter

func stderrWriter() (code int, err error) {
	var output *os.File
	if output, err = os.Open(`/var/log/redis.err`); nil != err {
		return
	}
	code, err = gex.Exec(`redis`, gex.Stderr(output))

	return
}
