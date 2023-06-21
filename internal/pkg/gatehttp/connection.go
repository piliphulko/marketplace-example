package gatehttp

import (
	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp/opt"
)

//s1 "github.com/piliphulko/marketplace-example/api/service-acct-aut"

var ( // requires connection
	ConnAA opt.ConnGrpc // s1.AccountAuthClient
)

const (
	connAA = iota + 1 // 1
)
