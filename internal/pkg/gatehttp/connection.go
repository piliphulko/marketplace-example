package gatehttp

import (
	s1 "github.com/piliphulko/marketplace-example/api/service-acct-aut"
)

var ( // requires connection
	ConnAA s1.AccountAuthClient
)

const (
	connAA = iota + 1 // 1
)
