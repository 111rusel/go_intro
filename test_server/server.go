package testserver

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(s Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- s.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			s.Cancel()
		}
	}
}

type StoreCtx interface {
	Fetch(ctx context.Context) (string, error)
}

func ServerCtx(s StoreCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data, err := s.Fetch(ctx)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Fprint(w, data)
	}
}
