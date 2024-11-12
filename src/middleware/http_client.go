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

func (this *HttpClient) createNewRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (this *HttpClient) Do(req *http.Request) (string, error) {
	resp, err := this.Client.Do(req)
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

func (this *HttpClient) Get(url string) (string, error) {
	req, err := this.createNewRequest(url)
	if err != nil {
		return "", err
	}

	return this.Do(req)
}
