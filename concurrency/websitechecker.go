package concurrency

// WebsiteChecker checks reachability of websites.
type WebsiteChecker func(string) bool

// CheckWebsites is the main entry function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
