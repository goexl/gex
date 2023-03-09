package gex

type trim struct {
	left  bool
	right bool
	all   bool
}

func newTrim() *trim {
	return &trim{
		left:  false,
		right: false,
		all:   false,
	}
}
