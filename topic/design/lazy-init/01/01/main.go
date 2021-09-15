package main

import "net/http"

type server struct{}

func (s *server) handleTemplate(files ...string) http.HandlerFunc {
	tpl, err := s.loadTemplate(files...)

	return func(w http.ResponseWriter, r *http.Request) {
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
