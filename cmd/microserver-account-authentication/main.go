package main

import (
	"log"
	"net"

	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	pb "github.com/piliphulko/marketplace-example/internal/service/microserver-account-authentication"
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
	pb.InitJWTSecret(viper.GetString("MICROSERVER-ACCOUNT-AUTHENTICATION.JWT_SECRET"))
}

func main() {
	err, logSync := logwriter.InitializeLoggerGRPC(
		&pb.LogGRPC, viper.GetString("MICROSERVER-ACCOUNT-AUTHENTICATION.LOG_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	defer logSync()

	lis, err := net.Listen(
		viper.GetString("MICROSERVER-ACCOUNT-AUTHENTICATION.NETWORK_SERVER"),
		viper.GetString("MICROSERVER-ACCOUNT-AUTHENTICATION.PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}

	grpclog.SetLoggerV2(pb.LogGRPC)

	var (
		grpcServer  = grpc.NewServer()
		microserver = pb.StartServer()
	)

	close, err := microserver.ConnPostrgresql(viper.GetString("POSTGRESQL.DATABASE_URL"))
	defer close()
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterServer(grpcServer, microserver)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
