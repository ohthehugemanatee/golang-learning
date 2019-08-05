package concurrency

// WebsiteChecker checks reachability of websites.
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites is the main entry function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		response := <-resultChannel
		results[response.string] = response.bool
	}
	return results
}
