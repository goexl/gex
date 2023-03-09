package collector

import (
	"fmt"

	"github.com/goexl/gex"
)

var _ = usage

type collector struct{}

func newCollector() *collector {
	return new(collector)
}

func (c *collector) Collect(line string, stream string) (err error) {
	fmt.Println(line)

	return
}

func usage() (int, error) {
	return gex.New("redis").Collect(newCollector()).Build().Exec()
}
