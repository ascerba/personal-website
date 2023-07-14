package main

import (
	"net/http"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if path[1] != "" {
		app.notFound(w)
		return
	} else {
		p, err := app.loadPosts("html/projects", 3)
		if err != nil {
			app.serverError(w, err)
			return
		}

		err = renderTemplate(w, "main/index", p)
		if err != nil {
			app.serverError(w, err)
			return
		}
	}
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "main/about", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) aggregate(w http.ResponseWriter, r *http.Request) {
	p, err := app.loadPosts("html"+strings.TrimSuffix(r.URL.Path, "/"), -1)
	if err != nil {
		app.notFound(w)
	}

	renderTemplate(w, "main/"+strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/"), "/"), p)
}

func (app *application) post(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if path[2] == "" {
		app.aggregate(w, r)
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
			return
		}
	}
}
