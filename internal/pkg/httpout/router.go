package httpout

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func RouterHTML() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", HandlerMainPage)
	r.Get("/customer/create", HandlerCustomerCreatePage)
	r.Post("/customer/create/send", HandlerCustomerCreateSendPage)

	return r
}
