package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/braveokafor/go-mail-api/handlers"
)

type HandlerTest struct {
	Name           string
	HandlerFunc    http.HandlerFunc
	ExpectedStatus int
	ExpectedBody   string
}

func TestHandlers(t *testing.T) {
	tests := []HandlerTest{
		{
			Name:           "HealthCheckHandler",
			HandlerFunc:    handlers.HealthCheck,
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   "{\"message\": \"ok\"}",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			server := httptest.NewServer(test.HandlerFunc)
			defer server.Close()

			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("error making request to server. Err: %v", err)
			}
			defer resp.Body.Close()

			// Assertions
			if resp.StatusCode != test.ExpectedStatus {
				t.Errorf("expected status %d; got %v", test.ExpectedStatus, resp.Status)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("error reading response body. Err: %v", err)
			}

			if test.ExpectedBody != string(body) {
				t.Errorf("expected response body to be %v; got %v", test.ExpectedBody, string(body))
			}
		})
	}
}
