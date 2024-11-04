package templating

import (
	"bytes"
	"testing"
)

func TestTemplating(t *testing.T) {
	var aPost = Post{
		Title: "Hello world",
		Body:  "This is a post",
		Tags: []string{
			"#go", "#tdd",
		},
	}
	t.Run("Render Title", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello world</h1>
<p>This is a post</p>
Tags: <ul><li>#go</li><li>#tdd</li></ul>`

		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})

}
