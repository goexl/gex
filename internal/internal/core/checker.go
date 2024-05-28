package core

type Checker interface {
	Check(line string) (checked bool, err error)
}
