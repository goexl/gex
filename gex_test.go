package gex_test

import (
	"testing"

	"github.com/goexl/gex"
	"github.com/goexl/gox/args"
)

func TestContains(t *testing.T) {
	t.Parallel()
	args := args.New().Build().Add("www.163.com").Build()
	_, err := gex.New("ping").Args(args).Checker().Contains("Ping statistics for").Async().Build().Exec()
	if nil != err {
		t.FailNow()
	}
}
