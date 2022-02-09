package gex_test

import (
	`testing`

	`github.com/storezhang/gex`
)

func TestExecWithCheckerSuccess(t *testing.T) {
	_, err := gex.Exec(
		`ping`, gex.Args(`www.163.com`),
		gex.ContainsChecker(`Ping statistics for`), gex.Async(), gex.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}

func TestExecWithCheckerFailed(t *testing.T) {
	_, err := gex.Exec(
		`ping`, gex.Args(`www.163.com`),
		gex.ContainsChecker(`xxx`), gex.Async(), gex.Quiet(),
	)
	if nil != err {
		t.FailNow()
	}
}

func TestExecWithSync(t *testing.T) {
	_, err := gex.Exec(`ping`, gex.Args(`www.163.com`), gex.Sync())
	if nil != err {
		t.FailNow()
	}
}

func TestExecWithStringCollector(t *testing.T) {
	output := ``
	_, err := gex.Exec(
		`ping`, gex.Args(`www.163.com`),
		gex.Sync(), gex.Quiet(), gex.StringCollector(&output),
	)
	if nil != err || `` == output {
		t.FailNow()
	}
}
