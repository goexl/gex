package gex

type lineCounter struct {
	total   *int64
	options *counterOptions
}

func newLineCounter(total *int64, options *counterOptions) *lineCounter {
	return &lineCounter{
		total:   total,
		options: options,
	}
}

func (c *lineCounter) Count(line string, ot OutputType) (err error) {
	if `` == line || OutputTypeAny != c.options.typ && ot != c.options.typ {
		return
	}
	*c.total++

	return
}
