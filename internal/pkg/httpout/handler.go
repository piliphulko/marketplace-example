package httpout

import (
	"html/template"
	"net/http"
)

type Into struct {
	Countries []string
	Cities    []string
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
	if html.Execute(w, struct{ GoodsARRAY []Goods }{
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
	}) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
