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

	r.Get("/belarus/marketplace", HandlerMarketplacePage)
	r.Get("/poland/marketplace", HandlerMarketplacePage)
	r.Get("/ukraine/marketplace", HandlerMarketplacePage)
	r.Post("/country/marketplace/send", HandlerMarketplaceSend)
	r.Get("/error", nil)

	r.Get("/", HandlerСleanPage("../../html/main-page.html"))
	r.Get("/customer/create", HandlerCustomerCreatePage)
	r.Post("/customer/create/send", HandlerCreateAccountSendPage("../../html/response-create.html", CUSTOMER))

	r.Get("/customer/authorization", HandlerСleanPage("../../html/customer-authorization.html"))
	r.Post("/customer/authorization/send", nil)

	r.Get("/partner", HandlerСleanPage("../../html/main-partner-page.html"))

	r.Get("/vendor/create", HandlerСleanPage("../../html/vendor-create.html"))
	r.Post("/vendor/create/send", HandlerCreateAccountSendPage("../../html/response-create.html", VENDOR))

	r.Get("/vendor/authorization", HandlerСleanPage("../../html/vendor-authorization.html"))
	r.Post("/vendor/authorization/send", nil)

	r.Get("/warehouse/create", HandlerWarehouseCreatePage)
	r.Post("/warehouse/create/send", HandlerCreateAccountSendPage("../../html/response-create.html", WAREHOUSE))

	r.Get("/warehouse/authorization", HandlerСleanPage("../../html/warehouse-authorization.html"))
	r.Post("/warehouse/authorization/send", nil)

	r.Get("/{login_customer}/marketplace", nil)
	r.Post("/{login_customer}/marketplace/send", nil)

	r.Get("/{login_customer}/home", nil)
	r.Post("/{login_customer}/home/order/repeal/{order_uuid}", nil)
	r.Post("/{login_customer}/home/order/confirm/{order_uuid}", nil)
	r.Post("/{login_customer}/home/order/pay/{order_uuid}", nil)

	r.Get("/{login_customer}/home/delivery/confirm", nil)

	r.Get("/{login_customer}/home/wallet", nil)

	r.Get("/{login_customer}/home/wallet/replenishment", nil)
	r.Post("/{login_customer}/home/wallet/replenishment/send", nil)

	return r
}
