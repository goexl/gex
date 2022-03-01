package collector

import (
	`github.com/goexl/gex`
)

var _ = filename

func filename() (int, error) {
	return gex.Exec(`redis`, gex.FilenameCollector(`/var/log/redis.out`))
}
