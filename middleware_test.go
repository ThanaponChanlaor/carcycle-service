package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func corsOriginFor(t *testing.T, allowedOrigins, requestOrigin string) string {
	t.Helper()
	t.Setenv("ALLOWED_ORIGINS", allowedOrigins)
	handler := withCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	req := httptest.NewRequest(http.MethodGet, "/api/cars", nil)
	req.Header.Set("Origin", requestOrigin)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec.Header().Get("Access-Control-Allow-Origin")
}

func TestCORSOpenWhenUnset(t *testing.T) {
	if got := corsOriginFor(t, "", "https://anything.example"); got != "*" {
		t.Fatalf("got %q, want *", got)
	}
}

func TestCORSAllowsListedOrigin(t *testing.T) {
	got := corsOriginFor(t, "https://carcycle.netlify.app/, http://localhost:3000", "https://carcycle.netlify.app")
	if got != "https://carcycle.netlify.app" {
		t.Fatalf("got %q, want the Netlify origin", got)
	}
}

func TestCORSRejectsUnlistedOrigin(t *testing.T) {
	if got := corsOriginFor(t, "https://carcycle.netlify.app", "https://evil.example"); got != "" {
		t.Fatalf("got %q, want no Access-Control-Allow-Origin header", got)
	}
}
