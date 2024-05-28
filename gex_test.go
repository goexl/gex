package gex_test

import (
	"testing"

	"github.com/goexl/args"
	"github.com/goexl/gex"
)

func TestContains(t *testing.T) {
	t.Parallel()
	arguments := args.New().Build().Add("www.163.com").Build()
	_, err := gex.New("ping").Arguments(arguments).Check().Contains("Ping statistics for").Build().Build().Async().Echo().Build().Exec()
	if nil != err {
		t.FailNow()
	}
}
