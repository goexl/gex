package gex

type checker interface {
	check(all string, line string) (checked bool, err error)
}
