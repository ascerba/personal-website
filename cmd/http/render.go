package main

import (
	"net/http"
	"strings"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmplPath string, p *Posts) (err error) {
	t, err := template.ParseFiles("html/master.tmpl.html", "html/"+tmplPath+".tmpl.html")
	if err != nil {
		return err
	}

	splitPath := strings.Split(tmplPath, "/")

	data := make(map[string]interface{})

	// If were loading the index, set page to 'Index' and pass through all posts.
	// Otherwise, set page to 'Projects' and pass through the first post (should only be one
	// coming in)
	if splitPath[0] == "index" {
		data["Page"] = "Index"
		data["Posts"] = p
	} else {
		data["Page"] = "Project"
		data["Post"] = p.Contents[0]
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
