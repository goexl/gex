package core

const (
	OperatorUnknown Operator = iota
	OperatorAnd
	OperatorOr
)

type Operator uint8
