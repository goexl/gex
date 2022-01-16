package gex_test

import (
	`testing`

	`github.com/storezhang/gex`
)

func TestRun(t *testing.T) {
	_, err := gex.Run(
		`ping`,
		gex.Args(`www.163.com`),
		gex.ContainsChecker(`Ping statistics for1`), gex.Async(), gex.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}
