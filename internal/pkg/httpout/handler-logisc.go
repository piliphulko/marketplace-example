package httpout

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

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
	redirectAnswer := RedirectAnswer{}
	defer r.Body.Close()
	if err := TakeRedirectAnswerFromURL(r, &redirectAnswer); err != nil {
		cancelCtxError(err)
	}
	if err := optHTTP.HTML.Execute(&buf, struct {
		RedirectAnswer RedirectAnswer
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
		RedirectAnswer: redirectAnswer,
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
	defer r.Body.Close()
	redirectAnswer := RedirectAnswer{}
	if err := TakeRedirectAnswerFromURL(r, &redirectAnswer); err != nil {
		cancelCtxError(err)
	}
	if err := optHTTP.HTML.Execute(&buf, struct {
		RedirectAnswer RedirectAnswer
		LoginWarehouse string
		GoodsARRAY     []struct {
			NameVendor string
			NameGoods  string
		}
	}{
		RedirectAnswer: redirectAnswer,
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

func warehouseHomeWalletPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	login_warehouse := chi.URLParam(r, "login_warehouse")
	_ = login_warehouse
	if err := optHTTP.HTML.Execute(&buf, struct {
		WalletMoneyAvailable float64
		WalletMoneyBlocked   float64
		CommissionPercentage float64
	}{
		WalletMoneyAvailable: 125.32,
		WalletMoneyBlocked:   4446.2,
		CommissionPercentage: 0.09 * 100,
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func handlerReceivingGoodsSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	var (
		login_warehouse = chi.URLParam(r, "login_warehouse")
		name_vendor     = r.FormValue("name_vendor")
		name_goods      = r.FormValue("name_goods")
		amount_goods    = r.FormValue("amount_goods")
	)

	fmt.Println(login_warehouse, name_vendor, name_goods, amount_goods)

	optHTTP.OkRedirectPath = strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_warehouse}", login_warehouse)
	optHTTP.ErrRedirectPath = strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_warehouse}", login_warehouse)

	if err := JSON.NewEncoder(&buf).Encode(RedirectAnswer{
		Ok:     true,
		OkInfo: "Goods was added successfully",
	}); err != nil {
		cancelCtxError(err)
		//ch <- []byte(strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_warehouse}", login_warehouse))
		return
	}

	ch <- buf.Bytes()
	//ch <- []byte(strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_warehouse}", login_warehouse))
}

func handlerWarehouseHomeChange(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	login_warehouse := chi.URLParam(r, "login_warehouse")
	if err := optHTTP.HTML.Execute(&buf, struct {
		LoginWarehouse       string
		NameWarehouse        string
		InfoWarehouse        string
		CommissionPercentage float64
	}{
		LoginWarehouse:       login_warehouse,
		NameWarehouse:        "test",
		InfoWarehouse:        "NOTHING",
		CommissionPercentage: 0.09 * 100,
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerWarehouseHomeChangeSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	var (
		login_warehouseURL = chi.URLParam(r, "login_warehouse")

		login_warehouse        = r.FormValue("login_warehouse")
		password_warehouse     = r.FormValue("password_warehouse")
		login_warehouse_new    = r.FormValue("login_warehouse_new")
		password_warehouse_new = r.FormValue("password_warehouse_new")
		commission_percentage  = r.FormValue("commission_percentage")
		info_warehouse         = r.FormValue("info_warehouse")
	)
	optHTTP.OkRedirectPath = strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_warehouse}", login_warehouse)
	optHTTP.ErrRedirectPath = strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_warehouse}", login_warehouse)

	fmt.Println(login_warehouseURL, login_warehouse, password_warehouse, login_warehouse_new, password_warehouse_new, commission_percentage, info_warehouse)

	if err := JSON.NewEncoder(&buf).Encode(RedirectAnswer{
		Ok:     true,
		OkInfo: "Account update completed successfully",
	}); err != nil {
		cancelCtxError(err)
		//ch <- []byte(strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_warehouse}", login_warehouse))
		return
	}
	ch <- buf.Bytes()
	//ch <- []byte(strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_warehouse}", login_warehouse))
}

func handlerWarehouseDeliveryConfirmSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	var (
		login_warehouse = chi.URLParam(r, "login_warehouse")
		login_customer  = chi.URLParam(r, "login_customer")
		order_uuid      = chi.URLParam(r, "order_uuid")
	)
	fmt.Println(login_warehouse, login_customer, order_uuid)

	optHTTP.OkRedirectPath = strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_warehouse}", login_warehouse)
	optHTTP.ErrRedirectPath = strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_warehouse}", login_warehouse)

	if err := JSON.NewEncoder(&buf).Encode(RedirectAnswer{
		Ok:     true,
		OkInfo: "order confirmed: " + order_uuid + "| login: " + login_customer,
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func vendorHomePage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	login_vendor := chi.URLParam(r, "login_vendor")
	redirectAnswer := RedirectAnswer{}
	defer r.Body.Close()
	if err := TakeRedirectAnswerFromURL(r, &redirectAnswer); err != nil {
		cancelCtxError(err)
	}
	if err := optHTTP.HTML.Execute(&buf, struct {
		RedirectAnswer RedirectAnswer
		LoginVendor    string
		WalletMoney    float64
		GoodsARRAY     []struct {
			NameWarehouse        string
			Country              string
			City                 string
			NameGoods            string
			TypeGoods            string
			AmountGoodsAvailable int
			AmountGoodsBlocked   int
			AmountGoodsDefect    int
			PriceGoods           float64
		}
	}{
		RedirectAnswer: redirectAnswer,
		LoginVendor:    login_vendor,
		WalletMoney:    13432,
		GoodsARRAY: []struct {
			NameWarehouse        string
			Country              string
			City                 string
			NameGoods            string
			TypeGoods            string
			AmountGoodsAvailable int
			AmountGoodsBlocked   int
			AmountGoodsDefect    int
			PriceGoods           float64
		}{
			{
				NameWarehouse:        "ssz",
				Country:              "BELARUS",
				City:                 "MINSK",
				NameGoods:            "ssz",
				TypeGoods:            "dscx",
				AmountGoodsAvailable: 13,
				AmountGoodsBlocked:   13,
				AmountGoodsDefect:    10,
				PriceGoods:           50.5,
			},
			{
				NameWarehouse:        "hgfdx",
				Country:              "BELARUS",
				City:                 "MINSK",
				NameGoods:            "ssz",
				TypeGoods:            "dscx",
				AmountGoodsAvailable: 13,
				AmountGoodsBlocked:   13,
				AmountGoodsDefect:    10,
				PriceGoods:           50.5,
			},
			{
				NameWarehouse:        "gfds",
				Country:              "BELARUS",
				City:                 "MINSK",
				NameGoods:            "ssz",
				TypeGoods:            "dscx",
				AmountGoodsAvailable: 13,
				AmountGoodsBlocked:   13,
				AmountGoodsDefect:    10,
				PriceGoods:           50.5,
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func handlerVendorHomeChange(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	login_warehouse := chi.URLParam(r, "login_vendor")
	if err := optHTTP.HTML.Execute(&buf, struct {
		LoginVendor string
		NameVendor  string
	}{
		LoginVendor: login_warehouse,
		NameVendor:  "gfdszcv",
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerVendorHomeChangeSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	var (
		login_vendorURL = chi.URLParam(r, "login_vendor")

		login_vendor        = r.FormValue("login_vendor")
		password_vendor     = r.FormValue("password_vendor")
		login_vendor_new    = r.FormValue("login_vendor_new")
		password_vendor_new = r.FormValue("password_vendor_new")
		name_vendor         = r.FormValue("name_vendor")
	)
	optHTTP.OkRedirectPath = strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_vendor}", login_vendorURL)
	optHTTP.ErrRedirectPath = strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_vendor}", login_vendorURL)

	fmt.Println(login_vendorURL, login_vendor, password_vendor, login_vendor_new, password_vendor_new, name_vendor)

	if err := JSON.NewEncoder(&buf).Encode(RedirectAnswer{
		Ok:     true,
		OkInfo: "Account update completed successfully",
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func vendorGoodsPricePage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	login_vendor := chi.URLParam(r, "login_vendor")
	redirectAnswer := RedirectAnswer{}
	defer r.Body.Close()
	if err := TakeRedirectAnswerFromURL(r, &redirectAnswer); err != nil {
		cancelCtxError(err)
	}
	if err := optHTTP.HTML.Execute(&buf, struct {
		RedirectAnswer RedirectAnswer
		LoginVendor    string
		GoodsListARRAY struct {
			NameGoods  []string
			TypeGoods  []string
			Country    []string
			SalesModel []string
		}
		GoodsARRAY []struct {
			NameGoods  string
			TypeGoods  string
			Country    string
			PriceGoods float64
			SalesModel string
		}
	}{
		RedirectAnswer: redirectAnswer,
		LoginVendor:    login_vendor,
		GoodsListARRAY: struct {
			NameGoods  []string
			TypeGoods  []string
			Country    []string
			SalesModel []string
		}{
			NameGoods:  []string{"NameGoods", "ds"},
			TypeGoods:  []string{"TypeGoods", "ds"},
			Country:    []string{"Country", "ds"},
			SalesModel: []string{"SalesModel", "ds"},
		},
		GoodsARRAY: []struct {
			NameGoods  string
			TypeGoods  string
			Country    string
			PriceGoods float64
			SalesModel string
		}{
			{
				NameGoods:  "dsz",
				TypeGoods:  "bv",
				Country:    "bvcx",
				PriceGoods: 10.1,
				SalesModel: "lifo",
			},
			{
				NameGoods:  "dsz",
				TypeGoods:  "bv",
				Country:    "bvcx",
				PriceGoods: 10.1,
				SalesModel: "lifo",
			},
			{
				NameGoods:  "dsz",
				TypeGoods:  "bv",
				Country:    "gfdxz",
				PriceGoods: 1560.1,
				SalesModel: "lifo",
			},
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func handlerVendorHomeGoodsPriceChangeSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	var (
		login_vendorURL = chi.URLParam(r, "login_vendor")

		name_goods   = r.FormValue("name_goods")
		country      = r.FormValue("country")
		sales_model  = r.FormValue("sales_model")
		change_price = r.FormValue("change_price")
	)
	optHTTP.OkRedirectPath = strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_vendor}", login_vendorURL)
	optHTTP.ErrRedirectPath = strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_vendor}", login_vendorURL)

	fmt.Println(login_vendorURL, name_goods, country, sales_model, change_price)

	if err := JSON.NewEncoder(&buf).Encode(RedirectAnswer{
		Ok:     true,
		OkInfo: "Update completed successfully",
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func handlerVendorHomeGoodsPriceAdditionSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	var (
		login_vendorURL = chi.URLParam(r, "login_vendor")

		name_goods   = r.FormValue("name_goods")
		country      = r.FormValue("country")
		sales_model  = r.FormValue("sales_model")
		change_price = r.FormValue("change_price")
	)
	optHTTP.OkRedirectPath = strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_vendor}", login_vendorURL)
	optHTTP.ErrRedirectPath = strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_vendor}", login_vendorURL)

	fmt.Println(login_vendorURL, name_goods, country, sales_model, change_price)

	if err := JSON.NewEncoder(&buf).Encode(RedirectAnswer{
		Ok:     true,
		OkInfo: "Addition completed successfully",
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func handlerVendorHomeGoodsPriceCreateSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *OptionsHTTP, r *http.Request, ch chan []byte) {
	buf := bytes.Buffer{}
	var (
		login_vendorURL = chi.URLParam(r, "login_vendor")

		name_goods = r.FormValue("name_goods")
		type_goods = r.FormValue("type_goods")
		info_goods = r.FormValue("info_goods")
	)
	optHTTP.OkRedirectPath = strings.ReplaceAll(optHTTP.OkRedirectPath, "{login_vendor}", login_vendorURL)
	optHTTP.ErrRedirectPath = strings.ReplaceAll(optHTTP.ErrRedirectPath, "{login_vendor}", login_vendorURL)

	fmt.Println(login_vendorURL, name_goods, type_goods, info_goods)

	if err := JSON.NewEncoder(&buf).Encode(RedirectAnswer{
		Ok:     true,
		OkInfo: "Creation completed successfully",
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}
