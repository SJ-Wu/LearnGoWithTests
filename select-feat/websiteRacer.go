package select_feat

import (
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) string {
	duration1 := measureResponseTime(url1)
	duration2 := measureResponseTime(url2)

	if duration1 < duration2 {
		return url1
	}
	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
