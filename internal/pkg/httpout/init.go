package httpout

import (
	"errors"
	"html/template"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

var (
	LogHTTP  *zap.Logger
	TempHTML *template.Template = template.New("html")
	JSON                        = jsoniter.ConfigCompatibleWithStandardLibrary
)

var (
	ErrReportedErrorNotList = errors.New("the reported error for the client does not match the list of possible errors")
	ErrNoClientError        = errors.New("no client error")
	ErrSpiderMan            = errors.New("SPIDER MAN")
)
