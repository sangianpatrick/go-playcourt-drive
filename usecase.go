package drive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Connect will connect to playcourt drive server and claim the session id.
func (c *Client) Connect(ctx context.Context) (sessionID string, err error) {
	url := fmt.Sprintf("%s%s", c.host, loginEndpoint)

	ctx, err = c.hook.BeforeProcess(ctx, url)
	if err != nil {
		return
	}

	defer c.hook.AfterProcess(ctx)

	creds := map[string]string{
		"username": c.username,
		"password": c.password,
	}

	credsBuff, _ := json.Marshal(creds)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(credsBuff))
	resp, err := c.ar.Do(ctx, req)

	if err := c.wrapError(err); err != nil {
		return sessionID, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return sessionID, ErrInvalidCredentials
	}

	responseBodyBuff, _ := ioutil.ReadAll(resp.Body)
	sessionID = string(responseBodyBuff)
	return
}

func (c *Client) wrapError(err error) error {
	if err == nil {
		return nil
	}

	if err == context.DeadlineExceeded {
		return err
	}

	return ErrUnexpected
}
