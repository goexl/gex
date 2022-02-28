package exec

type (
	checkerOption interface {
		applyChecker(options *checkerOptions)
	}

	checkerOptions struct {
		cache bool
	}
)

func defaultCheckerOptions() *checkerOptions {
	return &checkerOptions{
		cache: false,
	}
}
