package main

import (
	"log"
	"net/http"

	"github.com/piliphulko/marketplace-example/internal/pkg/accountaut"
	"github.com/piliphulko/marketplace-example/internal/pkg/httpout"
	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("../../config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := logwriter.InitializeLogger(&httpout.LogHTTP,
		viper.GetString("SERVER_HTTP_SEND_HTML.LOG_FILE.INFO_LEVEL"),
		viper.GetString("SERVER_HTTP_SEND_HTML.LOG_FILE.ERROR_LEVEL"),
		viper.GetString("SERVER_HTTP_SEND_HTML.LOG_FILE.PANIC_LEVEL")); err != nil {
		log.Fatal(err)
	}
	httpout.FillTempHTMLfromDir(viper.GetString("SERVER_HTTP_SEND_HTML.HTML_DIR"))

}

func main() {
	var (
		err         error
		closeConnAA func()
	)
	err, closeConnAA, httpout.ConnServerAA = accountaut.TakeConn(viper.GetString("SERVER_GRPC_ACCOUNT_AUT.PORT"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		closeConnAA()
		httpout.LogHTTP.Sync()
	}()

	log.Fatal(
		http.ListenAndServe(viper.GetString("SERVER_HTTP_SEND_HTML.PORT"), httpout.RouterHTML()),
	)
}
