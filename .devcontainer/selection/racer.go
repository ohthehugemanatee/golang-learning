package selection

import (
	"errors"
	"net/http"
	"time"
)

// ErrTimeout for when all sites fail to return quickly
var ErrTimeout = errors.New("All sites timed out")

// DefaultTimeout is basically hard coded, but this way we can run tests on a lower timeout.
var DefaultTimeout = 10 * time.Second

// RaceWebsites compares the response time of 2 urls
func RaceWebsites(urls [2]string) (string, error) {
	select {
	case <-raceWebsite(urls[0]):
		return urls[0], nil
	case <-raceWebsite(urls[1]):
		return urls[1], nil
	case <-time.After(DefaultTimeout):
		return "", ErrTimeout
	}
}

func raceWebsite(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
