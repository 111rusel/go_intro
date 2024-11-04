package templating

import (
	"html/template"
	"io"
)

const postTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>{{.Title}}</h1>
<p>{{.Body}}</p>
Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>
</body>
</html>`

type Post struct {
	Title string
	Body  string
	Tags  []string
}

func Render(w io.Writer, p Post) error {
	templ, err := template.New("Blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	err = templ.Execute(w, p)
	if err != nil {
		return err
	}
	return nil
	// _, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	// if err != nil {
	// 	return err
	// }

	// _, err = fmt.Fprintf(w, "\n<p>%s</p>", p.Body)
	// if err != nil {
	// 	return err
	// }

	// _, err = fmt.Fprintf(w, "\nTags: <ul>")
	// if err != nil {
	// 	return err
	// }

	// for i := range p.Tags {
	// 	_, err = fmt.Fprintf(w, "<li>%s</li>", p.Tags[i])
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// _, err = fmt.Fprintf(w, "</ul>")
	// if err != nil {
	// 	return err
	// }
	// return nil
}
