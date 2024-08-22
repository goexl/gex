package param

type Line struct {
	Stream string
	Total  *int
}

func NewLine(total *int) *Line {
	return &Line{
		Total: total,
	}
}
