package param

type Environment struct {
	Environments []string
	System       bool
}

func NewEnvironment() *Environment {
	return &Environment{
		Environments: make([]string, 0, 1),
		System:       true,
	}
}
