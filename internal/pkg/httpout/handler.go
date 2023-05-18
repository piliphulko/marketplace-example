package httpout

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

type Into struct {
	Countries       []string
	Cities          []string
	LoginCustomer   string
	CountryCustomer string
	CountryCity     string
}

type Goods struct {
	NameWarehouse string
	NameVendor    string
	TypeGoods     string
	NameGoods     string
	InfoGoods     string
	PriceGoods    string
	AmountGoods   string
}

type Order struct {
	NameWarehouse string
	NameVendor    string
	NameGoods     string
	AmountGoods   float64
	PriceGoods    float64
	Totalcost     float64
}

func HandlerCustomerCreatePage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("../../html/customer-create.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if html.Execute(w, Into{
		Countries: []string{
			"BELARUS",
			"POLAND",
			"UKRAINE",
		},
		Cities: []string{
			"MINSK",
			"WARSAW",
			"KYIV",
		},
	}) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerWarehouseCreatePage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("../../html/warehouse-create.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if html.Execute(w, Into{
		Countries: []string{
			"BELARUS",
			"POLAND",
			"UKRAINE",
		},
		Cities: []string{
			"MINSK",
			"WARSAW",
			"KYIV",
		},
	}) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerMarketplacePage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("../../html/marketplace.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if html.Execute(w,
		struct {
			GoodsARRAY  []Goods
			BoolOrders  bool
			OrderUuid   string
			OrdersARRAY []Order
		}{
			GoodsARRAY: []Goods{
				{
					NameWarehouse: "1",
					NameVendor:    "1",
					TypeGoods:     "1",
					NameGoods:     "1",
					InfoGoods:     "1",
					PriceGoods:    "1",
					AmountGoods:   "1",
				},
				{
					NameWarehouse: "2",
					NameVendor:    "2",
					TypeGoods:     "2",
					NameGoods:     "2",
					InfoGoods:     "2",
					PriceGoods:    "2",
					AmountGoods:   "2",
				},
			},
			BoolOrders: true,
			OrderUuid:  "qscdsad",
			OrdersARRAY: []Order{
				{
					NameWarehouse: "aa",
					NameVendor:    "aa",
					NameGoods:     "aa",
					AmountGoods:   1,
					PriceGoods:    1,
					Totalcost:     1,
				},
			},
		}) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerMarketplaceSend(w http.ResponseWriter, r *http.Request) {
	var (
		order_uuid     = r.FormValue("order_uuid")
		name_warehouse = r.FormValue("name_warehouse")
		name_vendor    = r.FormValue("name_vendor")
		name_goods     = r.FormValue("name_goods")
		amount_goods   = r.FormValue("amount_goods")
	)
	fmt.Println(order_uuid, name_warehouse, name_vendor, name_goods, amount_goods)
}

func HandlerCustomerHomePage(w http.ResponseWriter, r *http.Request) {
	login_customer := chi.URLParam(r, "login_customer")
	html, err := template.ParseFiles("../../html/customer-home.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u1, _ := uuid.NewV4()
	u2, _ := uuid.NewV4()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = html.Execute(w, struct {
		LoginCustomer              string
		WalletMoney                float64
		DeliveryConfirmARRAY       []string
		OrderStatusPlusOrdersARRAY []struct {
			OrderStatus string
			OrdersARRAY []struct {
				OrderUuid     string
				NameWarehouse string
				NameVendor    string
				NameGoods     string
				AmountGoods   int
				PriceGoods    float64
				Totalcost     float64
			}
		}
	}{
		LoginCustomer:        login_customer,
		WalletMoney:          1503.59,
		DeliveryConfirmARRAY: []string{u1.String(), u2.String()},
		OrderStatusPlusOrdersARRAY: []struct {
			OrderStatus string
			OrdersARRAY []struct {
				OrderUuid     string
				NameWarehouse string
				NameVendor    string
				NameGoods     string
				AmountGoods   int
				PriceGoods    float64
				Totalcost     float64
			}
		}{
			{
				OrderStatus: "unconfirmed order",
				OrdersARRAY: []struct {
					OrderUuid     string
					NameWarehouse string
					NameVendor    string
					NameGoods     string
					AmountGoods   int
					PriceGoods    float64
					Totalcost     float64
				}{
					{
						OrderUuid:     u1.String(),
						NameWarehouse: "111",
						NameVendor:    "111",
						NameGoods:     "111",
						AmountGoods:   5,
						PriceGoods:    10.3,
						Totalcost:     51.5,
					},
					{
						OrderUuid:     u2.String(),
						NameWarehouse: "111",
						NameVendor:    "111",
						NameGoods:     "111",
						AmountGoods:   1,
						PriceGoods:    10,
						Totalcost:     10,
					},
				},
			},
		},
	}); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerCustomerHomeWalletPage(w http.ResponseWriter, r *http.Request) {
	login_customer := chi.URLParam(r, "login_customer")
	html, err := template.ParseFiles("../../html/customer-home-wallet.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = html.Execute(w, struct {
		WalletMoney   float64
		LoginCustomer string
	}{
		WalletMoney:   100.56,
		LoginCustomer: login_customer,
	}); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerCustomerHomeWalletPromoSend(w http.ResponseWriter, r *http.Request) {
	var (
		promo_code = r.FormValue("promo_code")
	)
	fmt.Println(promo_code)
}

func HandlerCustomerHomeChangePage(w http.ResponseWriter, r *http.Request) {
	login_customer := chi.URLParam(r, "login_customer")
	html, err := template.ParseFiles("../../html/customer-home-change.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = html.Execute(w, Into{
		LoginCustomer:   login_customer,
		CountryCustomer: "BELARUS",
		CountryCity:     "MINSK",
		Countries: []string{
			"BELARUS",
			"POLAND",
			"UKRAINE",
		},
		Cities: []string{
			"MINSK",
			"WARSAW",
			"KYIV",
		},
	}); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerCustomerHomeChangeSend(w http.ResponseWriter, r *http.Request) {
	var (
		login        = r.FormValue("login")
		password_old = r.FormValue("password_old")
		login_new    = r.FormValue("login_new")
		password_new = r.FormValue("password_new")
		country      = r.FormValue("country")
		city         = r.FormValue("city")
	)
	fmt.Println(login, password_old, login_new, password_new, country, city)
}

func HandlerOrderConfirmSend(w http.ResponseWriter, r *http.Request) {
	login_customer := chi.URLParam(r, "login_customer")
	var (
		сonfirm_order_uuid = r.FormValue("сonfirm_order_uuid")
	)
	fmt.Println(сonfirm_order_uuid, login_customer)
}

func HandlerOrderPaySend(w http.ResponseWriter, r *http.Request) {
	login_customer := chi.URLParam(r, "login_customer")
	var (
		pay_order_uuid = r.FormValue("pay_order_uuid")
	)
	fmt.Println(pay_order_uuid, login_customer)
}

func HandlerOrderPepealSend(w http.ResponseWriter, r *http.Request) {
	login_customer := chi.URLParam(r, "login_customer")
	var (
		repeal_order_uuid = r.FormValue("repeal_order_uuid")
	)
	fmt.Println(repeal_order_uuid, login_customer)
}

func HandlerOrderDeliveryConfirmPage(w http.ResponseWriter, r *http.Request) {
	login_customer := chi.URLParam(r, "login_customer")
	_ = login_customer
	html, err := template.ParseFiles("../../html/customer-home-delivery-confirm.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = html.Execute(w, struct {
		OrderUuidARRAY []struct {
			OrderUuid     string
			NumberConfirm int
		}
	}{
		OrderUuidARRAY: []struct {
			OrderUuid     string
			NumberConfirm int
		}{
			{
				OrderUuid:     "ab",
				NumberConfirm: 5,
			},
			{
				OrderUuid:     "ac",
				NumberConfirm: 545,
			},
		},
	}); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerWarehouseHomePage(w http.ResponseWriter, r *http.Request) {
	login_warehouse := chi.URLParam(r, "login_warehouse")
	html, err := template.ParseFiles("../../html/warehouse-home.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u1, _ := uuid.NewV4()
	u2, _ := uuid.NewV4()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = html.Execute(w, struct {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerWarehouseInStockPage(w http.ResponseWriter, r *http.Request) {
	login_warehouse := chi.URLParam(r, "login_warehouse")
	_ = login_warehouse
	html, err := template.ParseFiles("../../html/warehouse-in-stock.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = html.Execute(w, struct {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerReceivingGoodsPage(w http.ResponseWriter, r *http.Request) {
	login_warehouse := chi.URLParam(r, "login_warehouse")
	_ = login_warehouse
	html, err := template.ParseFiles("../../html/warehouse-receiving-goods.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = html.Execute(w, struct {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandlerReceivingGoodsSend(w http.ResponseWriter, r *http.Request) {
	var (
		login_warehouse = chi.URLParam(r, "login_warehouse")
		name_vendor     = r.FormValue("name_vendor")
		name_goods      = r.FormValue("name_goods")
		amount_goods    = r.FormValue("amount_goods")
	)
	fmt.Println(login_warehouse, name_vendor, name_goods, amount_goods)
}

func HandlerWarehouseDeliveryConfirmSend(w http.ResponseWriter, r *http.Request) {
	var (
		login_warehouse   = chi.URLParam(r, "login_warehouse")
		login_customer    = chi.URLParam(r, "login_customer")
		order_uuid        = chi.URLParam(r, "order_uuid")
		confirmation_code = r.FormValue("confirmation_code")
	)
	fmt.Println(login_warehouse, login_customer, order_uuid, confirmation_code)
}

func HandlerWarehouseHomeWalletPage(w http.ResponseWriter, r *http.Request) {
	login_warehouse := chi.URLParam(r, "login_warehouse")
	_ = login_warehouse
	html, err := template.ParseFiles("../../html/warehouse-home-wallet.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = html.Execute(w, struct {
		WalletMoneyAvailable float64
		WalletMoneyBlocked   float64
		CommissionPercentage float64
	}{
		WalletMoneyAvailable: 125.32,
		WalletMoneyBlocked:   4446.2,
		CommissionPercentage: 0.09 * 100,
	}); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
