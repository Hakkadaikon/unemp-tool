package middleware

import (
	"testing"
)

func TestHttpClient_Do(t *testing.T) {
	t.Run("正常HTTPリクエスト", func(t *testing.T) {
		mockClient := &MockHttpClient{}
		mockClient.SetMockResponseOK()

		httpClient := &HttpClient{}
		httpClient.SetHttpClient(mockClient)

		req := mockClient.GetSampleHttpRequest()
		body, err := httpClient.Do(req)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if body != mockClient.GetResponseOKMessage() {
			t.Errorf(
				"expected body to be '%s', got %s",
				mockClient.GetResponseOKMessage(),
				body)
		}
	})

	t.Run("ネットワークエラー", func(t *testing.T) {
		mockClient := &MockHttpClient{}
		mockClient.SetMockNetworkError()

		httpClient := &HttpClient{}
		httpClient.SetHttpClient(mockClient)

		req := mockClient.GetSampleHttpRequest()
		_, err := httpClient.Do(req)
		if err == nil || err.Error() != mockClient.GetResponseNetworkErrorMessage() {
			t.Errorf(
				"expected error to be '%s', got %v",
				mockClient.GetResponseNetworkErrorMessage(),
				err)
		}
	})
}

func TestHttpClient_Get(t *testing.T) {
	t.Run("正常HTTPリクエスト", func(t *testing.T) {
		mockClient := &MockHttpClient{}
		mockClient.SetMockResponseOK()

		httpClient := &HttpClient{}
		httpClient.SetHttpClient(mockClient)

		body, err := httpClient.Get("http://example.com")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if body != mockClient.GetResponseOKMessage() {
			t.Errorf(
				"expected body to be '%s', got %s",
				mockClient.GetResponseOKMessage(),
				body)
		}
	})

	t.Run("不正URL", func(t *testing.T) {
		httpClient := &HttpClient{}
		_, err := httpClient.Get("http://\n")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
