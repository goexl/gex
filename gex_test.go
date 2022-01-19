package gex_test

import (
	`testing`

	`github.com/storezhang/gex`
)

func TestRunWithChecker(t *testing.T) {
	t.Parallel()
	_, err := gex.Run(
		`ping`, gex.Args(`www.163.com`),
		gex.ContainsChecker(`Ping statistics for`), gex.Async(), gex.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}

func TestRunWithSync(t *testing.T) {
	t.Parallel()
	_, err := gex.Run(`ping`, gex.Args(`www.163.com`), gex.Sync())
	if nil != err {
		t.FailNow()
	}
}

func TestRunWithStringCollector(t *testing.T) {
	t.Parallel()
	output := ``
	_, err := gex.Run(
		`ping`, gex.Args(`www.163.com`),
		gex.Sync(), gex.Quiet(), gex.StringCollector(&output),
	)
	if nil != err || `` == output {
		t.FailNow()
	}
}
