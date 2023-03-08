package gex

type checkerParams struct {
	cache bool
}

func newCheckerParams() *checkerParams {
	return &checkerParams{
		cache: false,
	}
}
