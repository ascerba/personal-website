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

	data := combineData(tmplPath, p)

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func combineData(path string, p *Posts) map[string]interface{} {
	splitPath := strings.Split(path, "/")

	data := make(map[string]interface{})

	if splitPath[0] == "main" {
		switch splitPath[1] {
		case "about":
			data["Page"] = splitPath[1]
		default:
			data["Page"] = splitPath[1]
			data["Posts"] = p
		}
	} else {
		data["Page"] = splitPath[0]
		data["Post"] = p.Contents[0]
	}

	return data
}
