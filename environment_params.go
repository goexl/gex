package gex

type environmentParams struct {
	environments []string
	system       bool
}

func newEnvironmentParams() *environmentParams {
	return &environmentParams{
		environments: make([]string, 0, 1),
		system:       true,
	}
}
