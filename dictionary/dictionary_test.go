package dictionary

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	t.Run("Search", func(t *testing.T) {
		dictionary := map[string]string{
			"test": "this is test",
		}
		want := "this is test"

		got := Search(dictionary, "test")

		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("Dictionary", func(t *testing.T) {
		d := Dictionary{
			"test": "this is test",
		}
		want := "this is test"

		got := d.Search("test")

		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("Dictionary", func(t *testing.T) {
		d := Dictionary{
			"test": "this is test",
		}
		want := ""

		got := d.Search("abcd")

		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})
}
