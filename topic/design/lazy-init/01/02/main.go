package main

import (
	"html/template"
	"net/http"
	"sync"
)

type server struct{}

func (s *server) handleTemplate(files ...string) http.HandlerFunc {
	var (
		init sync.Once
		tpl  *template.Template
		err  error
	)

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, err := s.loadTemplate(files...)
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	s.renderTemplate(tpl)
}

// ================================================================================

func main() {

	s := newServer()

	htt.HandleFunc("/", s.handleTemplate("layout.tplmhtml", "index.tpl.tml"))

}
