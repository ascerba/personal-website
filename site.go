package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type Post struct {
	Meta      []byte
	Thumbnail []byte
	Content   []byte
}

type Posts struct {
	Contents []*Post
}

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func loadPost(path string) (p *Posts, err error) {

	var section = 0
	var tmp *Post = new(Post)

	file, err := os.Open(path + ".html")
	if err != nil {
		file.Close()
		var tmpPosts []*Post
		tmpPosts = append(tmpPosts, tmp)
		return &Posts{Contents: tmpPosts}, err
	}

	scanner := bufio.NewScanner(file)

LineLoop:
	for scanner.Scan() {
		var line = scanner.Text()
		var lineByte = []byte(line + "\n")

		if line == "!@#" {
			section++
			goto LineLoop
		}

		switch section {
		case 0:
			tmp.Meta = append(tmp.Meta, lineByte...)
		case 1:
			tmp.Thumbnail = append(tmp.Thumbnail, lineByte...)
		case 2:
			tmp.Content = append(tmp.Content, lineByte...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	if err != nil {
		return nil, err
	}

	var tmpPosts []*Post
	tmpPosts = append(tmpPosts, tmp)

	return &Posts{Contents: tmpPosts}, nil
}

func loadPosts(location string, postCount int) (p *Posts, err error) {
	if postCount == 0 || postCount < -1 {
		return nil, os.ErrInvalid
	}

	var tmpPosts []*Post
	var depthCount = 0

	// TODO - Use less computationally heavy opperation for scalability
	err = filepath.Walk(location, func(path string, _ os.FileInfo, _ error) error {
		var directory, err = regexp.MatchString("^data/(projects|blog)$", path)
		if err != nil {
			return err
		}

		if directory {
			depthCount++
		} else if postCount == -1 || depthCount < postCount+1 {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			scanner := bufio.NewScanner(file)

			var section = 0
			var tmp *Post = new(Post)

		LineLoop:
			for scanner.Scan() {
				var line = scanner.Text()
				var lineByte = []byte(line + "\n")

				if line == "!@#" {
					section++
					goto LineLoop
				}

				switch section {
				case 0:
					tmp.Meta = append(tmp.Meta, lineByte...)
				case 1:
					tmp.Thumbnail = append(tmp.Thumbnail, lineByte...)
				case 2:
					tmp.Content = append(tmp.Content, lineByte...)
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			file.Close()

			tmpPosts = append(tmpPosts, tmp)
			depthCount++

		} else {
			depthCount++
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &Posts{Contents: tmpPosts}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Posts) {
	if tmpl == "post" {
		t, err := template.ParseFiles("templates/master.html", "templates/postHelp.tmpl")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, p.Contents[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		t, err := template.ParseFiles("templates/master.html", "templates/"+tmpl+".html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if path[1] != "" {
		http.NotFound(w, r)
	} else {
		p, err := loadPosts("data/projects", 3)
		if err != nil {
			http.NotFound(w, r)
		}

		renderTemplate(w, "index", p)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about", nil)
}

func agHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPosts("data"+strings.TrimSuffix(r.URL.Path, "/"), -1)
	if err != nil {
		http.NotFound(w, r)
	}

	renderTemplate(w, trimLeftChars(strings.TrimSuffix(r.URL.Path, "/"), 1), p)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if path[2] == "" {
		agHandler(w, r)
	} else {
		p, err := loadPost("data/" + path[1] + "/" + path[2])
		if p == nil || err != nil {
			http.NotFound(w, r)
		} else {
			renderTemplate(w, "post", p)
		}
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := os.ReadFile("data/www/" + r.URL.Path)

	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	fmt.Fprintf(w, "%s", body)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func httpsRedirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusMovedPermanently)
}

func main() {
	/* 	hostValid, _ := regexp.MatchString("^alexscerba.com$", r.URL.Path)
	   	if !hostValid {
	   		req, _ := http.NewRequest(r.Method, "alexscerba.com"+r.URL.Path, r.Body)
	   		redirect(w, req)
	   	} */

	fs := http.FileServer(http.Dir("./data/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/projects/", postHandler)
	http.HandleFunc("/blog/", postHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/", rootHandler)

	go http.ListenAndServe(":80", http.HandlerFunc(httpsRedirect))
	log.Fatal(http.ListenAndServe(":443", nil))

	//http.Handle("/", http.FileServer(http.Dir("/data/www")))
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
