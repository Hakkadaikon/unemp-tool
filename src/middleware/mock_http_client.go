package middleware

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

type MockHttpClient struct {
	response *http.Response
	err      error
}

func (this *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return this.response, this.err
}

func (this *MockHttpClient) SetMockResponse(response *http.Response, err error) {
	this.response = response
	this.err = err
}

func (this *MockHttpClient) GetResponseOKMessage() string {
	return "Hello, world!"
}

func (this *MockHttpClient) GetResponseNetworkErrorMessage() string {
	return "network error"
}

func (this *MockHttpClient) SetMockResponseOK() {
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(this.GetResponseOKMessage())),
	}
	this.SetMockResponse(mockResponse, nil)
}

func (this *MockHttpClient) SetMockNetworkError() {
	this.SetMockResponse(nil, errors.New(this.GetResponseNetworkErrorMessage()))
}

func (this *MockHttpClient) GetSampleHttpRequest() *http.Request {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	return req
}
