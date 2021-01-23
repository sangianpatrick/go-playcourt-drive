package drive

import (
	"context"
	"net/http"
	"net/url"
)

type defaultHook struct{}

// BeforeProcess initiates the span any request
func (r *defaultHook) BeforeProcess(ctx context.Context, cmd string) (context.Context, error) {
	return ctx, nil
}

// AfterProcess ends the initiated span from BeforeProcess
func (r *defaultHook) AfterProcess(ctx context.Context) error {
	return nil
}

// Hook is a playcourt drive hook collection of behavior.
type Hook interface {
	BeforeProcess(ctx context.Context, cmd string) (context.Context, error)
	AfterProcess(ctx context.Context) error
}

// Client is a concrete struct of playcourt drive client.
type Client struct {
	hook      Hook
	host      string
	sessionID string
	username  string
	password  string
	ar        *apiRequester
}

func isValidHost(host string) (valid bool) {
	u, err := url.Parse(host)
	if err != nil {
		return
	}

	switch u.Scheme {
	case "http":
		return true
	case "https":
		return true
	default:
		return
	}
}

func getMaxRetry(n int) (retries int) {
	defaultRetry := 3
	maxRetry := 7
	if n < defaultRetry || n > maxRetry {
		return defaultRetry
	}

	return n
}

func getMaxBackoff(ms int) (backoff int) {
	defaultBackoff := 10
	maxBackoff := 250
	if ms < defaultBackoff || ms > maxBackoff {
		return defaultBackoff
	}

	return ms
}

func getHTTPClient(c *http.Client) *http.Client {
	if c == nil {
		return http.DefaultClient
	}
	return c
}

// NewClient is a constructor of playcourt drive client.
func NewClient(config *Config) (c *Client, err error) {
	if config == nil {
		err = ErrEmptyConfig
		return
	}

	if config.Hook == nil {
		config.Hook = &defaultHook{}
	}

	if !isValidHost(config.Host) {
		err = ErrInvalidHost
		return
	}

	config.MaxRetry = getMaxRetry(config.MaxRetry)
	config.BackoffInMillis = getMaxBackoff(config.BackoffInMillis)
	config.CustomHTTPClient = getHTTPClient(config.CustomHTTPClient)

	ar := &apiRequester{
		hc:           config.CustomHTTPClient,
		maxRetry:     config.MaxRetry,
		retryBackoff: config.BackoffInMillis,
	}

	c = &Client{
		host:     config.Host,
		ar:       ar,
		password: config.Password,
		username: config.Username,
		hook:     config.Hook,
	}

	return
}
