package opt

import (
	s1 "github.com/piliphulko/marketplace-example/api/service-acct-aut"
)

type ConnGrpc interface {
	s1.AccountAuthClient
}
