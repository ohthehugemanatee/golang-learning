package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("basic functionality", func(t *testing.T) {
		fastServer := createDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()
		slowServer := createDelayedServer(2 * time.Millisecond)
		defer slowServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL
		urls := [2]string{fastURL, slowURL}

		want := fastURL
		got, _ := RaceWebsites(urls)

		if got != want {
			t.Errorf("Returned the wrong URL. Expected %q, got %q", want, got)
		}
	})
	t.Run("Timeout after 10 seconds", func(t *testing.T) {
		timeoutServer := createDelayedServer(2 * time.Second)
		defer timeoutServer.Close()
		timeoutServer2 := createDelayedServer(1 * time.Second)
		defer timeoutServer2.Close()
		url1 := timeoutServer.URL
		url2 := timeoutServer2.URL
		// Modify the timeout value so we don't have to wait during tests.``
		DefaultTimeout = 1 * time.Millisecond
		_, err := RaceWebsites([2]string{url1, url2})
		if err != ErrTimeout {
			t.Errorf("Expected a timeout error and got %q", err)
		}
	})
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
