package concurrency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func mockWebsiteChecker(url string) bool {
	time.Sleep(time.Millisecond * 20)
	return url != "error"
}

func TestCheckWebsites(t *testing.T) {
	t.Run("Check Websites", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"https://ya.ru",
			"error",
		}
		want := map[string]bool{
			"http://google.com": true,
			"https://ya.ru":     true,
			"error":             false,
		}

		got := CheckWebsites(mockWebsiteChecker, websites)
		require.Equal(t, want, got)
	})
}

func slowWebsitesChecker(_ string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsitesChecker, urls)
	}
}
