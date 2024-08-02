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

	t.Run("Add", func(t *testing.T) {
		d := Dictionary{}
		want := "this is test"

		d.Add("test", "this is test")

		got := d.Search("test")
		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("Add with error", func(t *testing.T) {
		d := Dictionary{
			"test": "this is test",
		}

		err := d.Add("test", "this is new test")
		if err == nil {
			t.Errorf("want error but got nil")
		}
	})

	t.Run("Update", func(t *testing.T) {
		d := Dictionary{
			"test": "this is test",
		}
		want := "this is new test"

		d.Update("test", "this is new test")

		got := d.Search("test")
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}

	})

	t.Run("Delete", func(t *testing.T) {
		d := Dictionary{"test": "this is test"}
		want := ""

		d.Delete("test")

		got := d.Search("test")
		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})
}
