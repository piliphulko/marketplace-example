package main

import (
	"log"
	"net/http"

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
		viper.GetString("SEND_HTML_SEVER.LOG_FILE.INFO_LEVEL"),
		viper.GetString("SEND_HTML_SEVER.LOG_FILE.ERROR_LEVEL"),
		viper.GetString("SEND_HTML_SEVER.LOG_FILE.PANIC_LEVEL")); err != nil {
		log.Fatal(err)
	}
	httpout.FillTempHTMLfromDir(viper.GetString("SEND_HTML_SEVER.HTML_DIR"))

}

func main() {
	log.Fatal(
		http.ListenAndServe(viper.GetString("SEND_HTML_SEVER.PORT"), httpout.RouterHTML()),
	)
}
