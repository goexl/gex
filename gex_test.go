package gex_test

import (
	`io/ioutil`
	`os`
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

func TestPwe(t *testing.T) {
	stdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	// 执行命令，该命令一定是错误的
	_, _ = gex.Exec(`ping`, gex.Args(`abc.c`))

	_ = writer.Close()
	out, _ := ioutil.ReadAll(reader)
	os.Stdout = stdout

	if `` == string(out) {
		t.FailNow()
	}
}
