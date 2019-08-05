package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	fastURL := fastServer.URL
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	slowURL := slowServer.URL
	urls := [2]string{fastURL, slowURL}
	want := fastURL
	got := RaceWebsites(urls)

	if got != want {
		t.Errorf("Returned the wrong URL. Expected %q, got %q", want, got)
	}
	slowServer.Close()
	fastServer.Close()
}
