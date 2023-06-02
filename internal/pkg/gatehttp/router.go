package gatehttp

import (
	"context"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp/opt"
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

	// PUBLIC ROUTER
	r.Group(func(r chi.Router) {
		r.Get("/marketplace/by", opt.NewOptionsHTTP().
			WithHTML(TempHTML, "marketplace-public.html").
			SetHeaderResponse(map[string]string{
				"Content-Type": "text/html; charset=utf-8"}).
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerMarketplacePublicBYPage))

		r.Get("/marketplace/pl", opt.NewOptionsHTTP().
			WithHTML(TempHTML, "marketplace-public.html").
			SetHeaderResponse(map[string]string{
				"Content-Type": "text/html; charset=utf-8"}).
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerMarketplacePublicPLPage))

		r.Get("/marketplace/ua", opt.NewOptionsHTTP().
			WithHTML(TempHTML, "marketplace-public.html").
			SetHeaderResponse(map[string]string{
				"Content-Type": "text/html; charset=utf-8"}).
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerMarketplacePublicUAPage))

		r.Get("/", opt.NewOptionsHTTP().
			WithHTML(TempHTML, "main-page.html").
			SetHeaderResponse(map[string]string{
				"Content-Type": "text/html; charset=utf-8"}).
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerCleanPage))

		r.Get("/partner", opt.NewOptionsHTTP().
			WithHTML(TempHTML, "main-partner-page.html").
			SetHeaderResponse(map[string]string{
				"Content-Type": "text/html; charset=utf-8"}).
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerCleanPage))
	})

	// CREATION AND AUTHORIZATION PUBLIC ROUTER
	r.Group(func(r chi.Router) {
		// CUSTOMER
		r.Get("/customer/create", opt.NewOptionsHTTP().
			ReceptionRedirectURL().
			WithHTML(TempHTML, "customer-create.html").
			SetHeaderResponse(map[string]string{
				"Content-Type": "text/html; charset=utf-8"}).
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerCleanPage))

		r.Post("/customer/create/send", opt.NewOptionsHTTP().
			URLSendRedirectOk("/customer/authorization").
			URLSendRedirectMistake("/customer/create").
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerCustomerCreateSend))

		r.Get("/customer/authorization", opt.NewOptionsHTTP().
			ReceptionRedirectURL().
			WithHTML(TempHTML, "customer-authorization.html").
			SetHeaderResponse(map[string]string{
				"Content-Type": "text/html; charset=utf-8"}).
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerAuthorizationPage))

		r.Post("/customer/authorization/send", opt.NewOptionsHTTP().
			URLSendRedirectOk("/{login_customer}/home").
			URLSendRedirectMistake("/customer/authorization").
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerCustomerAuthorizationSend))
	})

	// CUSTOMER PRIVAT ROUTER
	r.Group(func(r chi.Router) {

		r.Get("/{login_customer}/home", opt.NewOptionsHTTP().
			ReceptionRedirectURL().
			WithHTML(TempHTML, "customer-home.html").
			SetHeaderResponse(map[string]string{
				"Content-Type": "text/html; charset=utf-8"}).
			HandlerLogicsRun(context.Background(), time.Duration(5*time.Second), handlerCustomerHomePage))

	})

	return r
}
