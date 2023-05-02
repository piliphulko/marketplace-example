package httpout

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerMainPage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("../../html/main-page.html")
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

func HandlerCustomerCreatePage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("../../html/customer-create.html")
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

func HandlerCustomerCreateSendPage(w http.ResponseWriter, r *http.Request) {
	var (
		login        = r.FormValue("login")
		password_new = r.FormValue("password_new")
		country      = r.FormValue("country")
		city         = r.FormValue("city")
	)
	fmt.Println(login, password_new, country, city)
	html, err := template.ParseFiles("../../html/response-create.html")
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
