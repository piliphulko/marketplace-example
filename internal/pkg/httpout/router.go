package httpout

import (
	"context"
	"time"

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

	r.Get("/{login_customer}/home", HandlerCustomerHomePage)
	r.Get("/{login_customer}/home/change", HandlerCustomerHomeChangePage)
	r.Post("/{login_customer}/home/change/send", HandlerCustomerHomeChangeSend)

	r.Post("/{login_customer}/home/order/repeal/send", HandlerOrderPepealSend)
	r.Post("/{login_customer}/home/order/confirm/send", HandlerOrderConfirmSend)
	r.Post("/{login_customer}/home/order/pay/send", HandlerOrderPaySend)

	r.Get("/{login_customer}/home/delivery/confirm", nil)

	r.Get("/{login_customer}/home/wallet", HandlerCustomerHomeWalletPage)
	r.Post("/{login_customer}/home/wallet/promo/send", HandlerCustomerHomeWalletPromoSend)

	r.Get("/{login_warehouse}/warehouse/home", StartOptionsHTTP().
		WithHTML(TempHTML, "warehouse-home.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), time.Duration(5*time.Second), warehouseHomePage))

	r.Get("/{login_warehouse}/warehouse/home/change", StartOptionsHTTP().
		WithHTML(TempHTML, "warehouse-home-change.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerWarehouseHomeChange))

	r.Post("/{login_warehouse}/warehouse/home/change/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_warehouse}/warehouse/home").
		UseErrRedirectDataURL("/{login_warehouse}/warehouse/home").
		handlerRun(context.Background(), withTimeoutSecond(5), handlerReceivingGoodsSend))

	r.Get("/{login_warehouse}/warehouse/home/wallet", StartOptionsHTTP().
		WithHTML(TempHTML, "warehouse-home-wallet.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), warehouseHomeWalletPage))

	r.Get("/{login_warehouse}/in-stock/goods", StartOptionsHTTP().
		WithHTML(TempHTML, "warehouse-in-stock.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), testW))

	r.Get("/{login_warehouse}/receiving/goods", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "warehouse-receiving-goods.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), receivingGoodsPage))

	r.Post("/{login_warehouse}/receiving/goods/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_warehouse}/receiving/goods").
		UseErrRedirectDataURL("/{login_warehouse}/receiving/goods").
		handlerRun(context.Background(), withTimeoutSecond(5), handlerReceivingGoodsSend))

	r.Post("/{login_warehouse}/{login_customer}/{order_uuid}/delivery/confirm", HandlerWarehouseDeliveryConfirmSend)
	return r
}
