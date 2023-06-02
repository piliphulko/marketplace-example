package opt

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

var (
	LogZap *zap.Logger
	JSON   = jsoniter.ConfigCompatibleWithStandardLibrary
)
