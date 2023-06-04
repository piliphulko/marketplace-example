package main

import (
	"context"
	"fmt"
	"log"

	"github.com/piliphulko/marketplace-example/api/basic"
	pb "github.com/piliphulko/marketplace-example/api/microserver-account-authentication-client"
)

func main() {
	a, close, err := pb.ConnToMicroserverAccountAuthentication(":50051")
	defer close()
	if err != nil {
		log.Fatal(err)
	}
	reply, err := a.AutAccount(context.Background(), &basic.LoginPass{
		AccountChoice: &basic.LoginPass_CustomerLoginPass{
			CustomerLoginPass: &basic.CustomerAut{
				LoginCustomer:    "123",
				PasswortCustomer: "211",
			},
		},
	})
	fmt.Println(reply.Reply, err)
}
