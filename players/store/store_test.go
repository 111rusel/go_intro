package store

import "testing"

func TestStore(t *testing.T) {
	store := New()

	t.Run("TestStore", func(t *testing.T) {
		store.IncrementPlayersScore("Света")

		if store.playersScores["Света"] != 1 {
			t.Errorf("want score 1 but got %d", store.playersScores["Света"])
		}
	})

	t.Run("TestStore", func(t *testing.T) {
		store.IncrementPlayersScore("Света")

		if store.playersScores["Света"] != 2 {
			t.Errorf("want score 2 but got %d", store.playersScores["Света"])
		}
	})

	t.Run("TestStore", func(t *testing.T) {
		got := store.GetPlayersScore("Света")

		if got != 2 {
			t.Errorf("want score 2 but got %d", got)
		}
	})
}
