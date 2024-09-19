package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Run("Racer Test", func(t *testing.T) {
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL

		got, _ := Racer(slowURL, fastURL, time.Second)

		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("Racer Test 2", func(t *testing.T) {
		slowURL := slowServer.URL
		fastURL := slowServer.URL

		_, err := Racer(slowURL, fastURL, time.Millisecond*500)

		if err == nil {
			t.Errorf("want error but got nil")
		}
	})

	slowServer.Close()
	fastServer.Close()
}
