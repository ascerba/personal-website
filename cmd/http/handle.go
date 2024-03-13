package main

import (
	"net/http"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	path := strings.Split(r.URL.Path, "/")
	if path[1] != "" {
		app.notFound(w)
	} else {
		p, err := app.aggregate("html/projects")
		if err != nil {
			app.serverError(w, err)
		}

		err = renderTemplate(w, "index", p)
		if err != nil {
			app.serverError(w, err)
		}
	}
}

func (app *application) post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	path := strings.Split(r.URL.Path, "/")
	if len(path) > 4 {
		app.notFound(w)
	} else if len(path) == 4 && path[3] == "" {
		http.Redirect(w, r, "/"+path[1]+"/"+path[2], http.StatusFound)
	} else {
		post, err := app.readFile("html" + strings.TrimSuffix(r.URL.Path, "/") + ".tmpl.html")
		if err != nil {
			app.notFound(w)
			return
		}

		var posts []*Post
		posts = append(posts, post)
		p := &Posts{Contents: posts}

		err = renderTemplate(w, path[1]+"/"+path[2], p)
		if err != nil {
			app.serverError(w, err)
		}
	}
}
