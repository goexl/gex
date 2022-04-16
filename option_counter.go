package gex

var (
	_        = Counter
	_        = LineCounter
	_ option = (*optionCounter)(nil)
)

type optionCounter struct {
	counter counter
	err     error
}

// Counter 计数器
func Counter(counter counter) *optionCounter {
	return &optionCounter{
		counter: counter,
	}
}

// LineCounter 行计数器
func LineCounter(total *int64, opts ...counterOption) *optionCounter {
	return &optionCounter{
		counter: newLineCounter(total, newCounterOptions(opts...)),
	}
}

func (c *optionCounter) apply(options *options) (err error) {
	options.counters = append(options.counters, c.counter)

	return
}
