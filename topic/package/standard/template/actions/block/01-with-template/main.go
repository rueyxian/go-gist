package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

var (
	homeTmpl  *template.Template
	aboutTmpl *template.Template
)

// ============================================================

func main() {

	var err error
	var fs []string

	// ====================
	fs = []string{"base.html", "home.html"}
	homeTmpl, err = template.ParseFiles(fs...)
	if err != nil {
		panic(err)
	}
	// ====================
	fs = []string{"base.html", "about.html"}
	aboutTmpl, err = template.ParseFiles(fs...)
	if err != nil {
		panic(err)
	}
	// ====================
	http.HandleFunc("/", newHandler(homeTmpl))
	http.HandleFunc("/about", newHandler(aboutTmpl))

	http.ListenAndServe(":8080", nil)

}

func newHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
