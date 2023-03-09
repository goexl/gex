package gex

type check struct {
	operator operator
	checker  checker
}

func newCheck(operator operator, checker checker) *check {
	return &check{
		operator: operator,
		checker:  checker,
	}
}
