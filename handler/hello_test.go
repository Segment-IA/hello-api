package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	HelloHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status OK, got %d", rr.Code)
	}

	body := strings.TrimSpace(rr.Body.String())
	if body != "Hello, World" {
		t.Errorf("unexpected response body: got %q, want %q", body, "Hello, World")
	}
}
