package gex

type counter interface {
	Count(line string, steam string) error
}
