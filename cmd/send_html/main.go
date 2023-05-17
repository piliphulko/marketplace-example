package main

import (
	"log"
	"net/http"

	"github.com/piliphulko/marketplace-example/internal/pkg/httpout"
	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	"go.uber.org/zap"
)

func init() {
	if err := logwriter.InitializeLogger(httpout.LogHTTP, "infoHTTP.log", "errorHTTP.log", "panicHTTP.log"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var logger *zap.Logger
	if err := logwriter.InitializeLogger(logger, "infoHTTPtest.log", "errorHTTPtest.log", "panicHTTPtest.log"); err != nil {
		log.Fatal(err)
	}
	logger.Info("fdsz")

	httpout.FillTempHTMLfromDir("../../html")
	log.Fatal(
		http.ListenAndServe(":8080", httpout.RouterHTML()),
	)
}
