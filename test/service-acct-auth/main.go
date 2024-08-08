package main

import (
	"context"
	"fmt"
	"os"

	"github.com/piliphulko/marketplace-example/api/basic"
	s1 "github.com/piliphulko/marketplace-example/api/service-acct-aut"
)

func main() {
	s, close, err := s1.ConnToServiceAccountAuthentication(":30051")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer close()

	_, err = s.CreateAccount(context.Background(),
		&basic.CustomerNew{
			CustomerAut: &basic.CustomerAut{
				LoginCustomer:    "lol5567",
				PasswordCustomer: "12345678",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "BELARUS",
				CustomerCity:    "MINSK",
			},
		})
	fmt.Println(err)
}
