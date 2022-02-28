package exec

const (
	// OutputTypeStdout 输出流
	OutputTypeStdout OutputType = 1
	// OutputTypeStderr 错误流
	OutputTypeStderr OutputType = 2
	// OutputTypeAny 所有输出
	OutputTypeAny OutputType = 3
)

// OutputType 输出类型
type OutputType uint8
