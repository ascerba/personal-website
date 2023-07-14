package main

import (
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmplPath string, p *Posts) (err error) {
	t, err := template.ParseFiles("html/master.tmpl.html", "html/"+tmplPath+".tmpl.html")
	if err != nil {
		return err
	}

	err = t.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}
