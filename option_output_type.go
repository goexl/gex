package gex

var (
	_                 = CollectStdout
	_                 = CollectStderr
	_                 = CollectAny
	_ collectorOption = (*optionOutputType)(nil)
	_ counterOption   = (*optionOutputType)(nil)
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

func (ot *optionOutputType) applyCounter(options *counterOptions) {
	options.typ = ot.typ
}
