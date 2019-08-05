package selection

import (
	"net/http"
	"time"
)

// RaceWebsites compares the response time of 2 urls
func RaceWebsites(urls [2]string) string {
	speed1 := raceWebsite(urls[0])
	speed2 := raceWebsite(urls[1])
	if speed1 > speed2 {
		return urls[1]
	}
	return urls[0]
}

func raceWebsite(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
