package config

type Config struct {
	debug bool
}

func New(debug bool) *Config {
	return &Config{
		debug: debug,
	}
}

func (c *Config) Debug() bool {
	if c == nil {
		return false
	}
	return c.debug
}
