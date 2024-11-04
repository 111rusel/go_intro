package testserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write(b []byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

type SpyStoreCtx struct {
	response string
	t        *testing.T
}

func (s SpyStoreCtx) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				fmt.Println("Spy store got cancelled")
				return
			default:
				time.Sleep(time.Millisecond * 10)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyStore struct {
	response string
	canceled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(time.Millisecond * 100)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}

func TestServer(t *testing.T) {
	t.Run("serverTest", func(t *testing.T) {
		data := "hello World"
		server := Server(&SpyStore{response: data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("want %q got %q", response.Body.String(), data)
		}
	})

	t.Run("serverTest", func(t *testing.T) {
		data := "hello World"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(time.Millisecond*5, cancel)
		request = request.WithContext(ctx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !store.canceled {
			t.Errorf("request must be cancelled but done")
		}
	})

	t.Run("serverTest", func(t *testing.T) {
		data := "hello World"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("want %q got %q", response.Body.String(), data)
		}

		if store.canceled {
			t.Errorf("request must be done but cancelled")
		}

	})

	t.Run("server ctx test", func(t *testing.T) {
		data := "hello world"
		store := &SpyStoreCtx{response: data, t: t}
		svr := ServerCtx(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(time.Millisecond*5, cancel)
		request = request.WithContext(ctx)

		response := &SpyResponseWriter{}
		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("response should not have been written")
		}
	})
}
