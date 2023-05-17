package httpout

import (
	"html/template"

	"go.uber.org/zap"
)

var (
	LogHTTP  *zap.Logger
	TempHTML *template.Template = template.New("html")
)
