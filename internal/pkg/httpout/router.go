package httpout

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouterHTML() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("../../html/main_page.html")
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if html.Execute(w, nil) != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
	})
	return r
}
