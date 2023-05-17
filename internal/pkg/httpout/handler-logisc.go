package httpout

import (
	"bytes"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

func testW(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	if err := optHTTP.HTML.Execute(&buf, struct {
		GoodsARRAY []struct {
			ArrivalDate          string
			NameVendor           string
			TypeGoods            string
			NameGoods            string
			InfoGoods            string
			PriceGoods           float64
			AmountGoodsAvailable int
			AmountGoodsBlocked   int
			AmountGoodsDefect    int
		}
	}{
		GoodsARRAY: []struct {
			ArrivalDate          string
			NameVendor           string
			TypeGoods            string
			NameGoods            string
			InfoGoods            string
			PriceGoods           float64
			AmountGoodsAvailable int
			AmountGoodsBlocked   int
			AmountGoodsDefect    int
		}{
			{
				ArrivalDate:          "2023",
				NameVendor:           "dsa",
				TypeGoods:            "fds",
				NameGoods:            "gfdsv",
				InfoGoods:            "",
				PriceGoods:           15.5,
				AmountGoodsAvailable: 10,
				AmountGoodsBlocked:   41,
				AmountGoodsDefect:    1,
			},
			{
				ArrivalDate:          "2023",
				NameVendor:           "dsas",
				TypeGoods:            "fds",
				NameGoods:            "gfdsfv",
				InfoGoods:            "",
				PriceGoods:           152.5,
				AmountGoodsAvailable: 103,
				AmountGoodsBlocked:   413,
				AmountGoodsDefect:    14,
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func warehouseHomePage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	login_warehouse := chi.URLParam(r, "login_warehouse")
	u1, _ := uuid.NewV4()
	u2, _ := uuid.NewV4()
	if err := optHTTP.HTML.Execute(&buf, struct {
		LoginWarehouse string
		WalletMoney    float64
		OrdersARRAY    []struct {
			LoginCustomer string
			OrderUuid     string
			NameVendor    string
			NameGoods     string
			AmountGoods   int
			PriceGoods    float64
			Totalcost     float64
		}
	}{
		LoginWarehouse: login_warehouse,
		WalletMoney:    13432,
		OrdersARRAY: []struct {
			LoginCustomer string
			OrderUuid     string
			NameVendor    string
			NameGoods     string
			AmountGoods   int
			PriceGoods    float64
			Totalcost     float64
		}{
			{
				LoginCustomer: "ssz",
				OrderUuid:     u1.String(),
				NameVendor:    "ssz",
				NameGoods:     "ssz",
				AmountGoods:   1,
				PriceGoods:    13,
				Totalcost:     13,
			},
			{
				LoginCustomer: "ssz",
				OrderUuid:     u1.String(),
				NameVendor:    "ssz",
				NameGoods:     "ssz",
				AmountGoods:   1,
				PriceGoods:    13,
				Totalcost:     13,
			},
			{
				LoginCustomer: "ssz",
				OrderUuid:     u2.String(),
				NameVendor:    "ssz",
				NameGoods:     "ssz",
				AmountGoods:   1,
				PriceGoods:    132.5,
				Totalcost:     132.5,
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func receivingGoodsPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	login_warehouse := chi.URLParam(r, "login_warehouse")
	_ = login_warehouse
	if err := optHTTP.HTML.Execute(&buf, struct {
		LoginWarehouse string
		GoodsARRAY     []struct {
			NameVendor string
			NameGoods  string
		}
	}{
		LoginWarehouse: login_warehouse,
		GoodsARRAY: []struct {
			NameVendor string
			NameGoods  string
		}{
			{
				NameVendor: "ssz",
				NameGoods:  "ssz",
			},
			{
				NameVendor: "ssdss",
				NameGoods:  "sssssz",
			},
			{
				NameVendor: "sgggsz",
				NameGoods:  "sszfff",
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}
