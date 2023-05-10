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
