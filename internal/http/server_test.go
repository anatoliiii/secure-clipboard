package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anatoliiii/secure-clipboard/internal/auth"
	"github.com/anatoliiii/secure-clipboard/internal/clipboard"
)

type stubAuthenticator struct{}

func (stubAuthenticator) Authenticate(_ context.Context, username, password string) (auth.User, error) {
	if username == "demo" && password == "demo" {
		return auth.User{Username: "demo", DisplayName: "Demo"}, nil
	}
	return auth.User{}, auth.ErrInvalidCredentials
}

func TestServer_TextClipboardDefaultScope(t *testing.T) {
	t.Parallel()

	store := clipboard.NewScopedStore()
	authenticator := stubAuthenticator{}
	server := NewServer(store, authenticator, nil)
	handler := server.Handler()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/clipboard/text", nil)
	req.SetBasicAuth("demo", "demo")
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for empty clipboard, got %d", resp.Code)
	}

	payload := map[string]string{"content": "hello"}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	req = httptest.NewRequest(http.MethodPost, "/api/v1/clipboard/text", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("demo", "demo")
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/clipboard/text", nil)
	req.SetBasicAuth("demo", "demo")
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.Code)
	}

	var result textClipboardResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if result.Content != "hello" {
		t.Fatalf("expected content hello, got %s", result.Content)
	}
	if result.Owner != "demo" {
		t.Fatalf("expected owner demo, got %s", result.Owner)
	}
}

func TestServer_TextClipboardScopedIsolation(t *testing.T) {
	t.Parallel()

	store := clipboard.NewScopedStore()
	authenticator := stubAuthenticator{}
	server := NewServer(store, authenticator, nil)
	handler := server.Handler()

	payload := map[string]string{"content": "first"}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/clipboard/text?scope=local:device-a", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("demo", "demo")
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.Code)
	}

	payload["content"] = "second"
	body, err = json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}
	req = httptest.NewRequest(http.MethodPost, "/api/v1/clipboard/text", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Scope", "local:device-b")
	req.SetBasicAuth("demo", "demo")
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/clipboard/text", nil)
	req.SetBasicAuth("demo", "demo")
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for default scope, got %d", resp.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/clipboard/text?scope=local:device-a", nil)
	req.SetBasicAuth("demo", "demo")
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.Code)
	}
	var resultA textClipboardResponse
	if err := json.NewDecoder(resp.Body).Decode(&resultA); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resultA.Content != "first" {
		t.Fatalf("expected content first, got %s", resultA.Content)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/clipboard/text", nil)
	req.Header.Set("Scope", "local:device-b")
	req.SetBasicAuth("demo", "demo")
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.Code)
	}
	var resultB textClipboardResponse
	if err := json.NewDecoder(resp.Body).Decode(&resultB); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resultB.Content != "second" {
		t.Fatalf("expected content second, got %s", resultB.Content)
	}
}
