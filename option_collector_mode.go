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

// CollectCache 缓存
func CollectCache() *optionCollectorMode {
	return &optionCollectorMode{
		mode: CollectorModeCache,
	}
}

// CollectDirect 直接输出
func CollectDirect() *optionCollectorMode {
	return &optionCollectorMode{
		mode: CollectorModeDirect,
	}
}

func (m *optionCollectorMode) applyCollector(options *collectorOptions) {
	options.mode = m.mode
}
