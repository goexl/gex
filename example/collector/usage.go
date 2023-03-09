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

func (c *collector) Collect(line string, ot gex.OutputType) (err error) {
	switch ot {
	case gex.OutputTypeStderr:
		fmt.Println(line)
	case gex.OutputTypeStdout:
		fmt.Println(line)
	}

	return
}

func usage() (int, error) {
	return gex.Exec(`redis`, gex.Collector(newCollector()))
}
