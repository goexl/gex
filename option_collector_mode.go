package gex

var (
	_                 = CollectStdout
	_                 = CollectStderr
	_                 = CollectAny
	_ collectorOption = (*optionCollectorMode)(nil)
)

type optionCollectorMode struct {
	mode CollectorMode
}

// CollectStdout 只收集输出流
func CollectStdout() *optionCollectorMode {
	return &optionCollectorMode{
		mode: CollectorModeStdout,
	}
}

// CollectStderr 只收集错误流
func CollectStderr() *optionCollectorMode {
	return &optionCollectorMode{
		mode: CollectorModeStderr,
	}
}

// CollectAny 收集所有流
func CollectAny() *optionCollectorMode {
	return &optionCollectorMode{
		mode: CollectorModeAny,
	}
}

func (m *optionCollectorMode) applyCollector(options *collectorOptions) {
	options.mode = m.mode
}
