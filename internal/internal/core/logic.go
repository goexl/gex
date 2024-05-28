package core

import (
	"github.com/goexl/gox"
)

type Logic struct {
	operator Operator
	checker  Checker

	_ gox.CannotCopy
}

func NewLogic(operator Operator, checker Checker) *Logic {
	return &Logic{
		operator: operator,
		checker:  checker,
	}
}

func (l *Logic) Check(prev bool, line string) (checked bool, err error) {
	checked = true
	if result, ce := l.checker.Check(line); nil != ce {
		err = ce
	} else if OperatorAnd == l.operator {
		checked = prev && result
	} else if OperatorOr == l.operator {
		checked = prev || result
	}

	return
}

func (l *Logic) Next(prev bool) bool {
	return OperatorAnd == l.operator && !prev || OperatorOr == l.operator && prev
}
