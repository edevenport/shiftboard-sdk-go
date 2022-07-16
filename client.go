package shiftboard

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Default HTTP client values
const (
	BaseURLV1   = "https://m.shiftboard.com/api/v1"
	HTTPTimeout = time.Minute * 5
)

// NewClient initalizes the REST API credentials
func NewClient(email string, password string) *Client {
	c := &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: HTTPTimeout,
		},
	}

	c.Auth.Email = email
	c.Auth.Password = password
	c.Auth.UseUUID = true

	return c
}

func (c *Client) sendRequest(req *http.Request, v *Response) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	if c.Auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.Auth.AccessToken)
	}

	if c.Cookies != nil {
		for _, c := range c.Cookies {
			req.AddCookie(c)
		}
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	c.Cookies = resp.Cookies()

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		log.Printf("HTTP request failed: %v", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
