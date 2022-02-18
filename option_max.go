package gex

var (
	_                 = Max
	_ option          = (*optionMax)(nil)
	_ collectorOption = (*optionMax)(nil)
)

type optionMax struct {
	max int
}

// Max 配置最大收集行数
func Max(max int) *optionMax {
	return &optionMax{
		max: max,
	}
}

func (m *optionMax) apply(options *options) (err error) {
	options.max = m.max

	return
}

func (m *optionMax) applyCollector(options *collectorOptions) {
	options.max = m.max
}
