package gex

type checker interface {
	check(line string) (checked bool, err error)
}
