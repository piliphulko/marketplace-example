package httpout

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	CUSTOMER = iota + 1
	VENDOR
	WAREHOUSE
)

func HandlerСleanPage(htmlPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles(htmlPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if html.Execute(w, nil) != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func HandlerCreateAccountSendPage(htmlPath string, t int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles(string(htmlPath))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		switch t {
		case CUSTOMER:
			var (
				login        = r.FormValue("login")
				password_new = r.FormValue("password_new")
				country      = r.FormValue("country")
				city         = r.FormValue("city")
			)
			fmt.Println(login, password_new, country, city)
		case VENDOR:
			var (
				login           = r.FormValue("login")
				password_vendor = r.FormValue("password_vendor")
				name_vendor     = r.FormValue("name_vendor")
			)
			fmt.Println(login, password_vendor, name_vendor)
		case WAREHOUSE:
			var (
				login             = r.FormValue("login")
				password_warehose = r.FormValue("password_warehose")
				name_warehouse    = r.FormValue("name_warehouse")
				country           = r.FormValue("country")
				city              = r.FormValue("city")
				info_warehouse    = r.FormValue("info_warehouse")
			)
			fmt.Println(login, password_warehose, name_warehouse, country, city, info_warehouse)
		}
		if html.Execute(w, nil) != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
