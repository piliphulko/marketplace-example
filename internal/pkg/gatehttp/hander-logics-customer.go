package gatehttp

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/piliphulko/marketplace-example/api/basic"
	pbAA "github.com/piliphulko/marketplace-example/api/service-acct-aut"
	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp/opt"
)

func handlerCustomerCreateSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf               = bytes.Buffer{}
		login_customer    = r.FormValue("login_customer")
		password_customer = r.FormValue("password_customer")

		country = r.FormValue("country")
		city    = r.FormValue("city")
	)

	fmt.Println(login_customer, password_customer, country, city)

	if err := opt.WriteRedirectAnswerInfoOk(&buf, "Сreated"); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerCustomerAuthorizationSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf               = bytes.Buffer{}
		login_customer    = r.FormValue("login_customer")
		password_customer = r.FormValue("password_customer")
	)
	stringJWT, err := optHTTP.TakeConnGrpc(connAA).AutAccount(ctx,
		pbAA.OneofLoginPass(basic.CustomerAut{
			LoginCustomer:    login_customer,
			PasswortCustomer: password_customer,
		}))

	if err != nil {
		opt.WriteRedirectAnswerInfoErr(&buf, HandlerErrConnAA(err, cancelCtxError))
	} else {
		opt.WriteRedirectAnswerInfoOk(&buf, stringJWT.StringJwt)
	}

	ch <- buf.Bytes()
}

func handlerCustomerHomePage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf            = bytes.Buffer{}
		login_customer = chi.URLParam(r, "login_customer")
		u1, _          = uuid.NewV4()
		u2, _          = uuid.NewV4()
	)

	defer r.Body.Close()
	redirecrAnswer, err := opt.TakeRedirectAnswerFromURL(r)
	if err != nil {
		cancelCtxError(err)
		return
	}
	cookieJWT, err := r.Cookie(opt.CookieNameJWT)
	if err != nil {
		cancelCtxError(err)
		return
	}
	fmt.Println(cookieJWT.Value)

	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		RedirectAnswer         interface{}
		LoginCustomer          string
		WalletMoney            float64
		UnconfirmedOrdersARRAY []struct {
			OrderUUID     string
			Location      string
			NameWarehouse string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			PriceGoods    float64
			AmountGoods   int
		}
		СonfirmedOrdersARRAY []struct {
			OrderUUID     string
			Location      string
			NameWarehouse string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			PriceGoods    float64
			AmountGoods   int
		}
		HistoryOrdersARRAY []struct {
			OrderUUID     string
			Date          string
			Location      string
			NameWarehouse string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			PriceGoods    float64
			AmountGoods   int
		}
	}{
		RedirectAnswer: redirecrAnswer,
		LoginCustomer:  login_customer,
		WalletMoney:    1503.59,
		UnconfirmedOrdersARRAY: []struct {
			OrderUUID     string
			Location      string
			NameWarehouse string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			PriceGoods    float64
			AmountGoods   int
		}{
			{
				OrderUUID:     u1.String(),
				Location:      "bfdxv",
				NameWarehouse: "gfc",
				NameVendor:    "fxvcx",
				TypeGoods:     "gcvc",
				NameGoods:     "vcdx",
				PriceGoods:    100.2,
				AmountGoods:   3,
			},
			{
				OrderUUID:     u1.String(),
				Location:      "bfdxv",
				NameWarehouse: "gfc",
				NameVendor:    "fxvcx",
				TypeGoods:     "gcvc",
				NameGoods:     "vcdx",
				PriceGoods:    100.2,
				AmountGoods:   3,
			},
			{
				OrderUUID:     u2.String(),
				Location:      "bfdxv",
				NameWarehouse: "gfc",
				NameVendor:    "fxvcx",
				TypeGoods:     "gcvc",
				NameGoods:     "vcdx",
				PriceGoods:    100.2,
				AmountGoods:   3,
			},
		},
		СonfirmedOrdersARRAY: []struct {
			OrderUUID     string
			Location      string
			NameWarehouse string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			PriceGoods    float64
			AmountGoods   int
		}{
			{
				OrderUUID:     u2.String(),
				Location:      "bfdxv",
				NameWarehouse: "gfc",
				NameVendor:    "fxvcx",
				TypeGoods:     "gcvc",
				NameGoods:     "vcdx",
				PriceGoods:    100.2,
				AmountGoods:   3,
			},
		},
		HistoryOrdersARRAY: []struct {
			OrderUUID     string
			Date          string
			Location      string
			NameWarehouse string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			PriceGoods    float64
			AmountGoods   int
		}{
			{
				OrderUUID:     u2.String(),
				Date:          "10.10.2010",
				Location:      "bfdxv",
				NameWarehouse: "gfc",
				NameVendor:    "fxvcx",
				TypeGoods:     "gcvc",
				NameGoods:     "vcdx",
				PriceGoods:    100.2,
				AmountGoods:   3,
			},
			{
				OrderUUID:     u1.String(),
				Date:          "10.10.2010",
				Location:      "bfdxv",
				NameWarehouse: "gfc",
				NameVendor:    "fxvcx",
				TypeGoods:     "gcvc",
				NameGoods:     "vcdx",
				PriceGoods:    100.2,
				AmountGoods:   3,
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}
