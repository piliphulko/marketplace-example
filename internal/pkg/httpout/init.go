package httpout

import (
	"html/template"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

var (
	redirectUrl = "/{redirect_json}"
	LogHTTP     *zap.Logger
	TempHTML    *template.Template = template.New("html")
	JSON                           = jsoniter.ConfigCompatibleWithStandardLibrary
)
