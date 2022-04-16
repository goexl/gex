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

func (c *lineCounter) Count(line string, typ OutputType) (err error) {
	if `` == line {
		return
	}
	*c.total++

	return
}
