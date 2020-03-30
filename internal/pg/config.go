package pg

import "github.com/go-kit/kit/log"

// ConfigOption configures the client.
type ConfigOption func(*Client)

func NewClient(options ...ConfigOption) *Client {
	c := Client{
		logger:         log.NewNopLogger(),
		maxConnections: defaultMaxConnections,
	}

	for _, opt := range options {
		opt(&c)
	}

	return &c
}

func WithLogger(l log.Logger) ConfigOption {
	return func(c *Client) {
		c.logger = l
	}
}

func WithMaxConnections(n int) ConfigOption {
	return func(c *Client) {
		c.maxConnections = n
	}
}
