package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastServer := createDelayedServer(0 * time.Millisecond)
	fastURL := fastServer.URL
	slowServer := createDelayedServer(2 * time.Millisecond)
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

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
