package selection

import (
	"net/http"
	"time"
)

// RaceWebsites compares the response time of 2 urls
func RaceWebsites(urls [2]string) string {
	startA := time.Now()
	http.Get(urls[0])
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(urls[1])
	bDuration := time.Since(startB)

	if aDuration > bDuration {
		return urls[1]
	}
	return urls[0]
}
