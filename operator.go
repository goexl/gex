package gex

const (
	operatorUnknown operator = iota
	operatorAnd
	operatorOr
)

type operator uint8
