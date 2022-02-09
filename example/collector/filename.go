package collector

import (
	`github.com/storezhang/gex`
)

var _ = filename

func filename() (int, error) {
	return gex.Exec(`redis`, gex.FilenameCollector(`/var/log/redis.out`))
}
