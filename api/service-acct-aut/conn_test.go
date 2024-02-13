package serviceacctauth

import (
	"testing"

	"github.com/piliphulko/marketplace-example/api/basic"
)

var blackРole *basic.AccountInfoChange

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blackРole = OneofAccount[basic.CustomerChange](basic.CustomerChange{
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "vcx",
				PasswortCustomer: "gfdszvcxsa",
			},
			CustomerAutOld: &basic.CustomerAut{
				LoginCustomer:    "vcx",
				PasswortCustomer: "gfdszvcxsa",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "gfds",
				CustomerCity:    "vcxz ",
			},
		})
	}
}

// BenchmarkOne-8   	 2771151	       440.0 ns/op	     376 B/op	       6 allocs/op
// BenchmarkOne-8   	 2744019	       434.8 ns/op	     376 B/op	       6 allocs/op

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blackРole = &basic.AccountInfoChange{
			AccountInfo: &basic.AccountInfoChange_CustomerChange{
				CustomerChange: &basic.CustomerChange{
					CustomerAutNew: &basic.CustomerAut{
						LoginCustomer:    "vcx",
						PasswortCustomer: "gfdszvcxsa",
					},
					CustomerAutOld: &basic.CustomerAut{
						LoginCustomer:    "vcx",
						PasswortCustomer: "gfdszvcxsa",
					},
					CustomerInfo: &basic.CustomerInfo{
						CustomerCountry: "gfds",
						CustomerCity:    "vcxz ",
					},
				},
			},
		}
	}
}

// BenchmarkSimple-8   	 2603131	       440.2 ns/op	     376 B/op	       6 allocs/op
// BenchmarkSimple-8   	 2760363	       435.9 ns/op	     376 B/op	       6 allocs/op
