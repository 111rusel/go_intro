package main

import (
	"fmt"
	"go_intro/files"
	"go_intro/templating"
	"log"
	"net/http"
	"os"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := files.NewPostsFromFS(".", os.DirFS("./posts"))
	if err != nil {
		fmt.Println(err)
		return
	}

	p := templating.Post{
		Title: posts[0].Title,
		Body:  posts[0].Body,
		Tags:  []string{"#go", "#tdd", "#learn"},
	}
	err = templating.Render(w, p)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
