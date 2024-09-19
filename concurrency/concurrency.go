package concurrency

type WebsiteChecker func(string) bool

type result struct {
	s string
	b bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChanel := make(chan result)
	for _, url := range urls {
		go func(url string) {
			resultChanel <- result{s: url, b: wc(url)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChanel
		results[r.s] = r.b
	}
	return results
}
