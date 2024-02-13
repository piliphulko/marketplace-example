package main

import (
	"log"
	"net"

	"github.com/piliphulko/marketplace-example/internal/pkg/f16"
	pb "github.com/piliphulko/marketplace-example/internal/service/service-data-customer"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile("../../config/config.yaml")
	//viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	pb.InitJWTSecret(viper.GetString("SERVICE-ACCT-AUTH.JWT_SECRET"))
}

func main() {

	lis, err := net.Listen(
		viper.GetString("SERVICE-DATA-CUSTOMER.NETWORK_SERVER"),
		viper.GetString("SERVICE-DATA-CUSTOMER.PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			f16.InterceptorCheckCtx,
			f16.IntrceptorHandlerErrors,
		),
	)
	server := pb.StartServer()

	close, err := server.ConnPostrgresql(viper.GetString("POSTGRESQL.DATABASE_URL"))
	defer close()
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
