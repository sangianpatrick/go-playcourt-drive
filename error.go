package drive

import (
	"fmt"
)

// Collection of playcourt drive errors.
var (
	ErrClosed             = fmt.Errorf("Connection is closed")
	ErrInvalidHost        = fmt.Errorf("Invalid host")
	ErrEmptyConfig        = fmt.Errorf("Empty configuration")
	ErrInvalidCredentials = fmt.Errorf("Invalid credentials")
	ErrTimeout            = fmt.Errorf("Request timeout")
	ErrUnexpected         = fmt.Errorf("Unexpected error")
)
