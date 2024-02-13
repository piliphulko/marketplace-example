package main

import (
	"log"
	"net/http"

	s1 "github.com/piliphulko/marketplace-example/api/service-acct-aut"
	s3 "github.com/piliphulko/marketplace-example/api/service-data-customer"
	s2 "github.com/piliphulko/marketplace-example/api/service-order-pay"

	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("../../config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := gatehttp.FillTempHTMLfromDir(viper.GetString("SERVICE-HTTP-SEND-HTML.HTML_DIR")); err != nil {
		log.Fatal(err)
	}
	jwt.InsertSecretForSignJWS(viper.GetString("SERVICE-ACCT-AUTH.JWT_SECRET"))
}

func main() {
	var (
		closeConnAA, closeConnDC, closeConnOP func()
		err                                   error
	)
	gatehttp.ConnAA, closeConnAA, err = s1.ConnToServiceAccountAuthentication(viper.GetString("SERVICE-ACCT-AUTH.PORT"))
	if err != nil {
		log.Fatal(err)
	}

	defer closeConnAA()
	gatehttp.ConnOP, closeConnOP, err = s2.ConnToServiceOrderPay(viper.GetString("SERVICE-ORDER-PAY.PORT"))
	if err != nil {
		log.Fatal(err)
	}
	defer closeConnOP()
	gatehttp.ConnDC, closeConnDC, err = s3.ConnToServiceDataCustomer(viper.GetString("SERVICE-DATA-CUSTOMER.PORT"))
	if err != nil {
		log.Fatal(err)
	}
	defer closeConnDC()

	log.Fatal(
		http.ListenAndServe(viper.GetString("SERVICE-HTTP-SEND-HTML.PORT"), gatehttp.RouterHTML()),
	)
}
