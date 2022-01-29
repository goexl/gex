package gex

const (
	// CollectorTypeStdout 输出流
	CollectorTypeStdout CollectorType = 1
	// CollectorTypeStderr 错误流
	CollectorTypeStderr CollectorType = 2
	// CollectorTypeAny 所有输出
	CollectorTypeAny CollectorType = 3
)

// CollectorType 收集类型
type CollectorType uint8
