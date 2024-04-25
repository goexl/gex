package gex_test

import (
	"testing"

	"github.com/goexl/args"
	"github.com/goexl/gex"
)

func TestContains(t *testing.T) {
	t.Parallel()
	args := args.New().Build().Add("www.163.com").Build()
	_, err := gex.New("ping").Args(args).Checker().Contains("Ping statistics for").Async().Echo().Build().Exec()
	if nil != err {
		t.FailNow()
	}
}
