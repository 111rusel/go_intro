package files

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type MockFsForError struct {
}

func (m *MockFsForError) Open(name string) (fs.File, error) {
	return nil, errors.New("error")
}

func TestNewPostsFromFS(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(`Title: Hi
Body: Abracadabra`)},
			"hello world 2.md": {Data: []byte(`Title: Hola
Body: Aaabraacadabra`)},
		}
		posts, err := NewPostsFromFS(".", fs)

		if err != nil {
			t.Errorf("got error: %s", err.Error())
		}
		if len(posts) != len(fs) {
			t.Errorf("want %d posts, got %d posts", len(fs), len(posts))
		}
	})

	t.Run("Test2", func(t *testing.T) {
		fs := &MockFsForError{}
		_, err := NewPostsFromFS(".", fs)

		if err == nil {
			t.Errorf("want error but got nil")
		}
	})

	t.Run("Test3", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(`Title: post 1
Body: hfhfh`)},
			"hello world2.md": {Data: []byte(`Title: post 2
Body: ffgg`)},
		}
		posts, err := NewPostsFromFS(".", fs)

		if err != nil {
			t.Errorf("got error: %s", err.Error())
		}
		got := posts[0]
		want := Post{
			Title: "post 1",
			Body:  "hfhfh",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("Test4", func(t *testing.T) {
		const postBody = `Title: post 1
Body: Hello woorld`
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(postBody)},
		}

		posts, err := NewPostsFromFS(".", fs)
		if err != nil {
			t.Errorf("error: %s", err.Error())
		}

		got := posts[0]
		want := Post{Title: "post 1", Body: "Hello woorld"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
}
