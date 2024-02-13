package gatehttp

import (
	s1 "github.com/piliphulko/marketplace-example/api/service-acct-aut"
	s2 "github.com/piliphulko/marketplace-example/api/service-data-customer"
	s3 "github.com/piliphulko/marketplace-example/api/service-order-pay"
)

var ( // requires connection
	ConnAA s1.AccountAuthClient
	ConnDC s2.DataCustomerClient
	ConnOP s3.OrderPayClient
)
