package shiftboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) ListSites() (*Response, error) {
	payload := new(bytes.Buffer)
	if err := json.NewEncoder(payload).Encode(c.Auth); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/sites", c.BaseURL)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	resp := Response{}
	err = c.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("ListSites failed: %v %v", resp.Message, resp.Error)
	}

	return &resp, nil
}

func (c *Client) Login(orgID string) (*Response, error) {
	baseURL := fmt.Sprintf("%s/login", c.BaseURL)

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	// Set query parameters
	q := req.URL.Query()
	q.Add("orgID", orgID)
	q.Add("user_privacy", "true")
	req.URL.RawQuery = q.Encode()

	resp := Response{}
	err = c.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("Login failed: %v %v", resp.Message, resp.Error)
	}

	c.Auth.AccessToken = resp.Data.AccessToken

	return &resp, nil
}

func (c *Client) ListShifts(startDate string, endDate string) (*Response, error) {
	baseURL := fmt.Sprintf("%s/shifts", c.BaseURL)

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set query parameters
	q := req.URL.Query()
	q.Add("start_date", startDate)
	q.Add("end_date", endDate)
	q.Add("batch", "1000")
	req.URL.RawQuery = q.Encode()

	resp := Response{}
	err = c.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("ListShifts failed: %v %v", resp.Message, resp.Error)
	}

	return &resp, nil
}
