package gatehttp

import (
	"bytes"
	"context"
	"net/http"

	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp/opt"
)

func handlerCleanPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}

	defer r.Body.Close()
	redirecrAnswer, err := opt.TakeRedirectAnswerFromURL(r)
	if err != nil {
		cancelCtxError(err)
		return
	}
	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		RedirectAnswer interface{}
	}{
		RedirectAnswer: redirecrAnswer,
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerAuthorizationPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf = bytes.Buffer{}
	)
	defer r.Body.Close()
	redirecrAnswer, err := opt.TakeRedirectAnswerFromURL(r)
	if err != nil {
		cancelCtxError(err)
		return
	}

	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		RedirectAnswer interface{}
	}{
		RedirectAnswer: redirecrAnswer,
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerMarketplacePublicBYPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}

	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		GoodsARRAY []struct {
			NameWarehouse string
			Location      string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			InfoGoods     string
			PriceGoods    float64
		}
	}{
		GoodsARRAY: []struct {
			NameWarehouse string
			Location      string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			InfoGoods     string
			PriceGoods    float64
		}{
			{
				NameWarehouse: "dsa",
				Location:      "dsa",
				NameVendor:    "dsa",
				TypeGoods:     "dsa",
				NameGoods:     "dsa",
				InfoGoods:     "dsa",
				PriceGoods:    10.5,
			},
			{
				NameWarehouse: "dsa",
				Location:      "dsa",
				NameVendor:    "dsa",
				TypeGoods:     "dsa",
				NameGoods:     "dsa",
				InfoGoods:     "dsa",
				PriceGoods:    10.5,
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerMarketplacePublicPLPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}

	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		GoodsARRAY []struct {
			NameWarehouse string
			Location      string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			InfoGoods     string
			PriceGoods    float64
		}
	}{
		GoodsARRAY: []struct {
			NameWarehouse string
			Location      string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			InfoGoods     string
			PriceGoods    float64
		}{
			{
				NameWarehouse: "dsa",
				Location:      "dsa",
				NameVendor:    "dsa",
				TypeGoods:     "dsa",
				NameGoods:     "dsa",
				InfoGoods:     "dsa",
				PriceGoods:    10.5,
			},
			{
				NameWarehouse: "dsa",
				Location:      "dsa",
				NameVendor:    "dsa",
				TypeGoods:     "dsa",
				NameGoods:     "dsa",
				InfoGoods:     "dsa",
				PriceGoods:    10.5,
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerMarketplacePublicUAPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}

	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		GoodsARRAY []struct {
			NameWarehouse string
			Location      string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			InfoGoods     string
			PriceGoods    float64
		}
	}{
		GoodsARRAY: []struct {
			NameWarehouse string
			Location      string
			NameVendor    string
			TypeGoods     string
			NameGoods     string
			InfoGoods     string
			PriceGoods    float64
		}{
			{
				NameWarehouse: "dsa",
				Location:      "dsa",
				NameVendor:    "dsa",
				TypeGoods:     "dsa",
				NameGoods:     "dsa",
				InfoGoods:     "dsa",
				PriceGoods:    10.5,
			},
			{
				NameWarehouse: "dsa",
				Location:      "dsa",
				NameVendor:    "dsa",
				TypeGoods:     "dsa",
				NameGoods:     "dsa",
				InfoGoods:     "dsa",
				PriceGoods:    10.5,
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}
