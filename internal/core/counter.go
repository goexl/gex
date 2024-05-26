package core

type Counter interface {
	Count(line string, steam string) error
}
