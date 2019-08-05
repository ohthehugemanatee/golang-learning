package selection

import (
	"fmt"
	"net/http"
	"time"
)

// RaceWebsites compares the response time of 2 urls
func RaceWebsites(urls [2]string) (string, error) {
	select {
	case <-raceWebsite(urls[0]):
		return urls[0], nil
	case <-raceWebsite(urls[1]):
		return urls[1], nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("All sites timed out")
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
