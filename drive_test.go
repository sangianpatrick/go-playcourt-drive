package drive_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	drive "github.com/sangianpatrick/go-playcourt-drive"
	"github.com/stretchr/testify/assert"
)

func TestNewClient_ErrEmptyConfig(t *testing.T) {
	c, err := drive.NewClient(nil)

	assert.Error(t, err)
	assert.Equal(t, drive.ErrEmptyConfig, err)
	assert.Nil(t, c)
}

func TestNewClient_Success(t *testing.T) {
	c, err := drive.NewClient(&drive.Config{
		BackoffInMillis:  100,
		CustomHTTPClient: http.DefaultClient,
		Host:             "http://drive.playcourt.test",
		MaxRetry:         5,
	})

	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewClient_ErrInvalidHost(t *testing.T) {
	t.Run("valid but not http or https", func(t *testing.T) {
		c, err := drive.NewClient(&drive.Config{
			Host: "mongodb://drive.playcourt.test",
		})

		assert.Error(t, err)
		assert.Nil(t, c)
	})

	t.Run("excatly invalid", func(t *testing.T) {
		c, err := drive.NewClient(&drive.Config{
			Host: ",,::",
		})

		assert.Error(t, err)
		assert.Nil(t, c)
	})
}

func TestNewClient_Success_WithoutHTTPClient_WithoutMaxRetry_WithoutBackoff_WithoutHook(t *testing.T) {
	c, err := drive.NewClient(&drive.Config{
		Host: "https://drive.playcourt.test",
	})

	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestConnect_ErrUnexpected(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMovedPermanently)
		io.WriteString(w, "301 Moved Permanently")
	}))

	c, _ := drive.NewClient(&drive.Config{
		BackoffInMillis:  2,
		CustomHTTPClient: http.DefaultClient,
		Host:             server.URL,
		MaxRetry:         3,
	})

	sessID, err := c.Connect(context.TODO())

	assert.Error(t, err)
	assert.Equal(t, drive.ErrUnexpected, err)
	assert.Equal(t, "", sessID)
}

func TestConnect_ErrContextDeadline(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "OK")
	}))

	c, _ := drive.NewClient(&drive.Config{
		BackoffInMillis:  2,
		CustomHTTPClient: http.DefaultClient,
		Host:             server.URL,
		MaxRetry:         3,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
	defer cancel()

	sessID, err := c.Connect(ctx)

	assert.Error(t, err)
	assert.Equal(t, context.DeadlineExceeded, err)
	assert.Equal(t, "", sessID)
}

func TestConnect_ErrInvalidCredentials(t *testing.T) {
	resultBody := "403 Forbidden"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, resultBody)
	}))

	c, _ := drive.NewClient(&drive.Config{
		BackoffInMillis:  2,
		CustomHTTPClient: http.DefaultClient,
		Host:             server.URL,
		MaxRetry:         3,
	})

	sessID, err := c.Connect(context.TODO())

	assert.Error(t, err)
	assert.Equal(t, drive.ErrInvalidCredentials, err)
	assert.Equal(t, "", sessID)
}

func TestConnect_Success(t *testing.T) {
	sessionIDMock := "let say that is a session id"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, sessionIDMock)
	}))

	c, _ := drive.NewClient(&drive.Config{
		BackoffInMillis:  2,
		CustomHTTPClient: http.DefaultClient,
		Host:             server.URL,
		MaxRetry:         3,
	})

	sessID, err := c.Connect(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, sessionIDMock, sessID)
}
