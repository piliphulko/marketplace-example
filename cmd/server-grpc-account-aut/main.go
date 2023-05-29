package main

import (
	"log"
	"net"

	"github.com/piliphulko/marketplace-example/internal/pkg/accountaut"
	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func init() {
	viper.SetConfigFile("../../config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err, logSync := logwriter.InitializeLoggerGRPC(
		&accountaut.LogGRPC, viper.GetString("SERVER_GRPC_ACCOUNT_AUT.LOG_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	defer logSync()

	lis, err := net.Listen(
		viper.GetString("SERVER_GRPC_ACCOUNT_AUT.NETWORK_SERVER"),
		viper.GetString("SERVER_GRPC_ACCOUNT_AUT.PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}

	grpclog.SetLoggerV2(accountaut.LogGRPC)

	var grpcServer = grpc.NewServer()
	accountaut.RegisterSever(grpcServer, accountaut.NewServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
