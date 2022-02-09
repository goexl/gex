package collector

import (
	`github.com/storezhang/gex`
)

var _ = options

func options() (code int, err error) {
	output := ``
	code, err = gex.Exec(`redis`, gex.StringCollector(
		&output,
		// 只收集100行数据
		gex.Max(100),
		// 只收集标准流出
		// gex.CollectStdout(),
		// 只收集标准错误
		// gex.CollectStderr(),
		// 收集所有数据
		gex.CollectAny(),
		// 直接收集
		gex.CollectDirect(),
		// 缓存起来，当命令执行完成后一起收集
		// gex.CollectCache(),
	))

	return
}
