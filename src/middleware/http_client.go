package middleware

import (
	"io"
	"net/http"
)

type HttpClient struct {
	Client HttpClientInterface
}

func NewHttpClient(client HttpClientInterface) *HttpClient {
	return &HttpClient{Client: client}
}

func (c *HttpClient) Get(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
