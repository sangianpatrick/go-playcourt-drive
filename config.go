package drive

import (
	"net/http"
)

// Config is a concrete struct of playcourt drive configuration.
type Config struct {
	Host             string
	Username         string
	Password         string
	CustomHTTPClient *http.Client
	BackoffInMillis  int
	MaxRetry         int
	Hook             Hook
}
