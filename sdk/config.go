package sdk

import "time"

var defaultEndpoint = "wakatime.com"

type Config struct {
	Scheme   string
	Endpoint string
	Timeout  time.Duration
}

func NewConfig() *Config {
	return &Config{
		Scheme:  SchemeHTTPS,
		Timeout: 20 * time.Second,
	}
}

func (c *Config) WithScheme(scheme string) *Config {
	c.Scheme = scheme
	return c
}

func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

func (c *Config) WithEndoint(endpoint string) *Config {
	c.Endpoint = endpoint
	return c
}
