package server

import (
	"go_intro/players/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *MockPlayerStore) GetPlayersScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *MockPlayerStore) IncrementPlayersScore(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayers(t *testing.T) {
	store := &MockPlayerStore{
		map[string]int{
			"Ruslan": 20,
			"Danis":  30,
		},
		[]string{},
	}

	server := &PlayerServer{
		store,
	}

	t.Run("Return Ruslan score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Ruslan", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "{\"score\": 20}"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if response.Code != http.StatusOK {
			t.Errorf("got %d want %d", response.Code, http.StatusOK)
		}
	})

	t.Run("Return Danis score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Danis", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "{\"score\": 30}"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if response.Code != http.StatusOK {
			t.Errorf("got %d want %d", response.Code, http.StatusOK)
		}
	})

	t.Run("Score Sveta", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Sveta", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSaveScore(t *testing.T) {
	store := &MockPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{
		store,
	}
	t.Run("Test Save", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Ruslan", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusAccepted {
			t.Errorf("got %d want %d", response.Code, http.StatusAccepted)
		}
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to increment score, but want 1", len(store.winCalls))
		}
	})
}

func TestFullIntegration(t *testing.T) {
	t.Run("Full Integration Test", func(t *testing.T) {
		st := store.New()
		srv := NewPlayerServer(st)

		request, _ := http.NewRequest(http.MethodPost, "/players/Ruslan", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)
		srv.ServeHTTP(response, request)
		srv.ServeHTTP(response, request)

		request, _ = http.NewRequest(http.MethodGet, "/players/Ruslan", nil)
		response = httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("want 200 got %d", response.Code)
		}

		if response.Body.String() != `{"score": 3}` {
			t.Errorf(`want {"score": 3} but got %q`, response.Body.String())
		}
	})
}
