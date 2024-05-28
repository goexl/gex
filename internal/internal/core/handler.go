package core

type Handler interface {
	Pid() int
	Kill() error
	Code() int
}
