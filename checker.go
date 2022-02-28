package exec

type checker interface {
	Check(line string) (checked bool, err error)
}
