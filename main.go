package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	fullchain = "/etc/letsencrypt/live/alexscerba.com/fullchain.pem"
	privkey   = "/etc/letsencrypt/live/alexscerba.com/privkey.pem"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func (app *application) httpsRedirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	app.infoLog.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusMovedPermanently)
}

func (app *application) wwwRedirect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.Host, "www.") {
			http.Redirect(w, r, "https://www."+r.Host+r.RequestURI, 302)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse() // required before flag is used

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/projects", app.post)
	mux.HandleFunc("/projects/", app.post)
	mux.HandleFunc("/blog", app.post)
	mux.HandleFunc("/blog/", app.post)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/about/", app.about)
	mux.HandleFunc("/", app.home)

	www := app.wwwRedirect(mux)

	if *addr == ":443" {
		infoLog.Printf("Starting TLS server on %s...\n", *addr)
		go http.ListenAndServe(":80", www)
		err := http.ListenAndServeTLS(*addr, fullchain, privkey, gzipHandler(www))
		log.Fatal(err)
	} else {
		infoLog.Printf("Starting server on %s...\n", *addr)
		err := http.ListenAndServe(*addr, gzipHandler(www))
		log.Fatal(err)
	}
}
