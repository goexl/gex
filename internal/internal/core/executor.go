package core

type Executor interface {
	Exec() (Handler, error)
}
