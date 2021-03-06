package collector

import (
	`os`

	`github.com/goexl/gex`
)

var _ = file

func file() (code int, err error) {
	var output *os.File
	if output, err = os.Open(`/var/log/redis.err`); nil != err {
		return
	}
	code, err = gex.Exec(`redis`, gex.FileCollector(output))

	return
}
