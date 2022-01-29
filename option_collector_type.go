package gex

var (
	_                 = CollectStdout
	_                 = CollectStderr
	_                 = CollectAny
	_ collectorOption = (*optionCollectorType)(nil)
)

type optionCollectorType struct {
	typ CollectorType
}

// CollectStdout 只收集输出流
func CollectStdout() *optionCollectorType {
	return &optionCollectorType{
		typ: CollectorTypeStdout,
	}
}

// CollectStderr 只收集错误流
func CollectStderr() *optionCollectorType {
	return &optionCollectorType{
		typ: CollectorTypeStderr,
	}
}

// CollectAny 收集所有流
func CollectAny() *optionCollectorType {
	return &optionCollectorType{
		typ: CollectorTypeAny,
	}
}

func (t *optionCollectorType) applyCollector(options *collectorOptions) {
	options.typ = t.typ
}
