package gex

import (
	"strings"
)

type trim struct {
	left  string
	right string
	all   string
	space bool
}

func newTrim() *trim {
	return &trim{
		left:  "",
		right: "",
		space: false,
	}
}

func (t *trim) process(line string) (final string) {
	final = line
	if "" != t.left {
		final = strings.TrimLeft(final, t.left)
	}
	if "" != t.right {
		final = strings.TrimRight(final, t.right)
	}
	if t.space {
		strings.TrimSpace(final)
	}

	return
}
