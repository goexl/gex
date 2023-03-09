package gex

type counterLines struct {
	total  *int
	params *counterParams
}

func newLineCounter(total *int, params *counterParams) *counterLines {
	return &counterLines{
		total:  total,
		params: params,
	}
}

func (cl *counterLines) Count(line string, stream string) (err error) {
	if "" != line && ("" == cl.params.stream || stream == cl.params.stream) {
		*cl.total++
	}

	return
}
