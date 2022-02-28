package exec

var (
	_                 = CollectStdout
	_                 = CollectStderr
	_                 = CollectAny
	_ collectorOption = (*optionOutputType)(nil)
)

type optionOutputType struct {
	typ OutputType
}

// CollectStdout 只收集输出流
func CollectStdout() *optionOutputType {
	return &optionOutputType{
		typ: OutputTypeStdout,
	}
}

// CollectStderr 只收集错误流
func CollectStderr() *optionOutputType {
	return &optionOutputType{
		typ: OutputTypeStderr,
	}
}

// CollectAny 收集所有流
func CollectAny() *optionOutputType {
	return &optionOutputType{
		typ: OutputTypeAny,
	}
}

func (ot *optionOutputType) applyCollector(options *collectorOptions) {
	options.typ = ot.typ
}
