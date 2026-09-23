package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClientUsesCache(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Write([]byte(`{"count":0,"results":[]}`))
	}))
	defer srv.Close()

	client := NewClient(5*time.Second, 5*time.Minute)
	url := srv.URL

	if _, err := client.ListLocations(&url); err != nil {
		t.Fatal(err)
	}
	if _, err := client.ListLocations(&url); err != nil {
		t.Fatal(err)
	}

	if calls != 1 {
		t.Errorf("expected 1 HTTP call, got %d", calls)
	}
}
