package httpout

import (
	"context"

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

	r.Get("/marketplace/by", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "marketplace-public.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerMarketplacePublicBYPage))

	r.Get("/marketplace/pl", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "marketplace-public.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerMarketplacePublicPLPage))

	r.Get("/marketplace/ua", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "marketplace-public.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerMarketplacePublicUAPage))

	r.Get("/", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "main-page.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCleanPage))

	r.Get("/customer/create", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "customer-create.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCleanPage))

	r.Post("/customer/create/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/customer/authorization").
		UseErrRedirectDataURL("/customer/create").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerCreateSend))

	r.Get("/customer/authorization", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "customer-authorization.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerAuthorizationPage))

	r.Get("/customer/authorization/send", StartOptionsHTTP().
		WithConnServerAccountAut(ConnServerAA).
		UseOkRedirectDataURL("/{login_customer}/home").
		UseErrRedirectDataURL("/customer/authorization").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerAuthorizationSend))

	r.Get("/partner", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "main-partner-page.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCleanPage))

	r.Get("/vendor/create", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "vendor-create.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCleanPage))

	r.Post("/vendor/create/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/vendor/authorization").
		UseErrRedirectDataURL("/vendor/create").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerVendorCreateSend))

	r.Get("/vendor/authorization", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "vendor-authorization.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerAuthorizationPage))

	r.Get("/vendor/authorization/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_vendor}/vendor/home").
		UseErrRedirectDataURL("/vendor/authorization").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerVendorAuthorizationSend))

	r.Get("/warehouse/create", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "warehouse-create.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCleanPage))

	r.Post("/warehouse/create/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/warehouse/authorization").
		UseErrRedirectDataURL("/warehouse/create").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerWarehouseCreateSend))

	r.Get("/warehouse/authorization", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "warehouse-authorization.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerAuthorizationPage))

	r.Get("/warehouse/authorization/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_warehouse}/warehouse/home").
		UseErrRedirectDataURL("/warehouse/authorization").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerWarehouseAuthorizationSend))

	r.Get("/{login_customer}/marketplace", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "marketplace.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerMarketplaceCustomerPage))

	r.Post("/{login_customer}/marketplace/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_customer}/marketplace").
		UseErrRedirectDataURL("/{login_customer}/marketplace").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerMarketplaceCustomerSend))

	r.Get("/{login_customer}/home", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "customer-home.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerHomePage))

	r.Get("/{login_customer}/home/change", StartOptionsHTTP().
		WithHTML(TempHTML, "customer-home-change.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerHomeChangePage))

	r.Post("/{login_customer}/home/change/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_customer}/home").
		UseErrRedirectDataURL("/{login_customer}/home").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerHomeChangeSend))

	r.Post("/{login_customer}/{order_uuid}/confirm/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_customer}/home").
		UseErrRedirectDataURL("/{login_customer}/home").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerHomeConfirmSend))

	r.Post("/{login_customer}/{order_uuid}/cancellation/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_customer}/home").
		UseErrRedirectDataURL("/{login_customer}/home").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerHomeCancellationSend))

	r.Get("/{login_customer}/{order_uuid}/receiving", StartOptionsHTTP().
		WithHTML(TempHTML, "customer-home-receiving.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerHomeReceivingPage))

	r.Get("/{login_customer}/home/wallet", StartOptionsHTTP().
		WithHTML(TempHTML, "customer-home-wallet.html").
		ReceptionRedirectURL().
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerHomeWalletPage))

	r.Post("/{login_customer}/home/wallet/promo/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_customer}/home/wallet").
		UseErrRedirectDataURL("/{login_customer}/home/wallet").
		handlerRun(context.Background(), withTimeoutSecond(5), handlerCustomerHomeWalletSend))

	r.Get("/{login_warehouse}/warehouse/home", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "warehouse-home.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), warehouseHomePage))

	r.Get("/{login_warehouse}/warehouse/home/change", StartOptionsHTTP().
		WithHTML(TempHTML, "warehouse-home-change.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerWarehouseHomeChange))

	r.Post("/{login_warehouse}/warehouse/home/change/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_warehouse}/warehouse/home").
		UseErrRedirectDataURL("/{login_warehouse}/warehouse/home").
		handlerRun(context.Background(), withTimeoutSecond(5), handlerWarehouseHomeChangeSend))

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
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerReceivingGoodsSend))

	r.Post("/{login_warehouse}/{login_customer}/{order_uuid}/delivery/confirm/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_warehouse}/warehouse/home").
		UseErrRedirectDataURL("/{login_warehouse}/warehouse/home").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerWarehouseDeliveryConfirmSend))

	r.Get("/{login_vendor}/vendor/home", StartOptionsHTTP().
		ReceptionRedirectURL().
		WithHTML(TempHTML, "vendor-home.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), vendorHomePage))

	r.Get("/{login_vendor}/vendor/home/change", StartOptionsHTTP().
		WithHTML(TempHTML, "vendor-home-change.html").
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerVendorHomeChange))

	r.Post("/{login_vendor}/vendor/home/change/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_vendor}/vendor/home").
		UseErrRedirectDataURL("/{login_vendor}/vendor/home").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerVendorHomeChangeSend))

	r.Get("/{login_vendor}/vendor/home/goods/price", StartOptionsHTTP().
		WithHTML(TempHTML, "vendor-home-goods-price.html").
		ReceptionRedirectURL().
		HeaderHTTPResponse(map[string]string{
			"Content-Type": "text/html; charset=utf-8"}).
		handlerRun(context.Background(), withTimeoutSecond(5), vendorGoodsPricePage))

	r.Post("/{login_vendor}/vendor/home/goods/price/change/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_vendor}/vendor/home/goods/price").
		UseErrRedirectDataURL("/{login_vendor}/vendor/home/goods/price").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerVendorHomeGoodsPriceChangeSend))

	r.Post("/{login_vendor}/vendor/home/goods/price/addition/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_vendor}/vendor/home/goods/price").
		UseErrRedirectDataURL("/{login_vendor}/vendor/home/goods/price").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerVendorHomeGoodsPriceAdditionSend))

	r.Post("/{login_vendor}/vendor/home/goods/price/create/send", StartOptionsHTTP().
		UseOkRedirectDataURL("/{login_vendor}/vendor/home/goods/price").
		UseErrRedirectDataURL("/{login_vendor}/vendor/home/goods/price").
		SetErrorClientList(ErrSpiderMan).
		handlerRun(context.Background(), withTimeoutSecond(5), handlerVendorHomeGoodsPriceCreateSend))

	return r
}
