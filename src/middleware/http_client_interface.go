package middleware

import (
	"net/http"
)

type HttpClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}
