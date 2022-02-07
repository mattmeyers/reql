package req

import (
	"bytes"
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c *Client) Do(req Request) (*http.Request, *http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080%s", req.Path)
	httpReq, err := http.NewRequest(req.Method, url, bytes.NewBufferString(req.Body))
	if err != nil {
		return nil, nil, err
	}

	res, err := c.client.Do(httpReq)
	if err != nil {
		return nil, nil, err
	}

	return httpReq, res, nil
}