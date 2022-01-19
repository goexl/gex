package gex

const (
	// CollectorModeStdout 输出流
	CollectorModeStdout CollectorMode = 1
	// CollectorModeStderr 错误流
	CollectorModeStderr CollectorMode = 2
	// CollectorModeAny 所有输出
	CollectorModeAny CollectorMode = 3
)

// CollectorMode 收集模式
type CollectorMode uint8
