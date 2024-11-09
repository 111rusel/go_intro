package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayersScore(name string) int
	IncrementPlayersScore(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{
		store: store,
	}
}

// GET /players/{name} - возвращает число побед игрока
// POST /players/{name} - записывает победу игрока

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.setScore(w, r)
	case http.MethodGet:
		s.showScore(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	score := s.store.GetPlayersScore(playerName)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "{\"score\": %d}", score)
}

func (s *PlayerServer) setScore(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	s.store.IncrementPlayersScore(playerName)
	w.WriteHeader(http.StatusAccepted)

}
