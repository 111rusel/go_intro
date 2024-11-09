package store

func New() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		playersScores: map[string]int{},
	}
}

type InMemoryPlayerStore struct {
	playersScores map[string]int
}

func (s *InMemoryPlayerStore) GetPlayersScore(name string) int {
	return s.playersScores[name]
}

func (s *InMemoryPlayerStore) IncrementPlayersScore(name string) {
	s.playersScores[name]++
}
