package infra

type config struct {
	maxLen int
}

type Option func(*config)

func WithMaxSize(len int) Option {
	return func(c *config) {
		c.maxLen = len
	}
}
