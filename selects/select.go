package selects

import (
	"errors"
	"net/http"
	"time"
)

func Racer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("ошибка")
	}
}

func ping(URL string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(URL)
		close(ch)
	}()
	return ch
}
