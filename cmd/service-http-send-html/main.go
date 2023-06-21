package main

import (
	"log"
	"net/http"

	serviceacctauth "github.com/piliphulko/marketplace-example/api/service-acct-aut"
	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp"
	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp/opt"
	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

var (
	logSync, closeConnAA func()
	err                  error
)

func init() {
	viper.SetConfigFile("../../config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if logSync, err = logwriter.InitZapLog(
		&opt.LogZap, viper.GetString("SERVICE-HTTP-SEND-HTML.LOG_FILE.ERROR_LEVEL"), zapcore.ErrorLevel,
	); err != nil {
		log.Fatal(err)
	}

	if err := gatehttp.FillTempHTMLfromDir(viper.GetString("SERVICE-HTTP-SEND-HTML.HTML_DIR")); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer logSync()
	gatehttp.ConnAA, closeConnAA, err = serviceacctauth.ConnToServiceAccountAuthentication(viper.GetString("SERVICE-ACCT-AUTH.PORT"))
	if err != nil {
		log.Fatal(err)
	}
	defer closeConnAA()

	log.Fatal(
		http.ListenAndServe(viper.GetString("SERVER_HTTP_SEND_HTML.PORT"), gatehttp.RouterHTML()),
	)
}
