package main

import (
	"log"
	"net/http"

	pbClient "github.com/piliphulko/marketplace-example/api/service-acct-aut"
	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp"
	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func init() {
	viper.SetConfigFile("../../config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	logSyng, err := logwriter.InitZapLog(&gatehttp.LogHTTP, viper.GetString("SERVICE-HTTP-SEND-HTML.LOG_FILE.INFO_LEVEL"), zapcore.InfoLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer logSyng()

	gatehttp.FillTempHTMLfromDir(viper.GetString("SERVICE-HTTP-SEND-HTML.HTML_DIR"))

}

func main() {
	var (
		err         error
		closeConnAA func()
	)
	gatehttp.ConnAA, closeConnAA, err = pbClient.ConnToServiceAccountAuthentication(":50051")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		closeConnAA()
		gatehttp.LogHTTP.Sync()
	}()

	log.Fatal(
		http.ListenAndServe(viper.GetString("SERVICE-HTTP-SEND-HTML.PORT"), gatehttp.RouterHTML()),
	)
}
