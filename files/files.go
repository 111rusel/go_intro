package files

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title string
	Body  string
}

func NewPostsFromFS(path string, myFS fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(myFS, path)
	if err != nil {
		return []Post{}, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(myFS, f)
		if err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(myFS fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := myFS.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	post, err := newPost(postFile)
	return post, err
}

func newPost(f io.Reader) (Post, error) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	titleLine := scanner.Text()
	scanner.Scan()
	bodyLine := scanner.Text()
	return Post{
		Title: titleLine[7:],
		Body:  bodyLine[6:],
	}, nil
}
