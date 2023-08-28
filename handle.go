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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	path := strings.Split(r.URL.Path, "/")
	pathLen := len(path)

	if pathLen == 3 && path[2] == "" {
		http.Redirect(w, r, "/about", http.StatusFound)
		return
	} else if pathLen == 3 && path[2] != "" || pathLen > 3 {
		app.notFound(w)
		return
	}
	err := renderTemplate(w, "main/about", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) aggregate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	p, err := app.loadPosts("html"+strings.TrimSuffix(r.URL.Path, "/"), -1)
	if err != nil {
		app.notFound(w)
	}

	renderTemplate(w, "main/"+strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/"), "/"), p)
}

func (app *application) post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	path := strings.Split(r.URL.Path, "/")
	if len(path) > 4 {
		app.notFound(w)
	} else if r.URL.Path == "/blog" || r.URL.Path == "/projects" { // Make a more encompasing change here
		app.aggregate(w, r)
	} else if path[2] == "" {
		http.Redirect(w, r, "/"+path[1], http.StatusFound)
		return
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
			return
		}
	}
}
