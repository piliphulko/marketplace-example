package main

import (
	"log"
	"net"

	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	pb "github.com/piliphulko/marketplace-example/internal/service/service-acct-aut"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func init() {
	viper.SetConfigFile("../../config/config.yaml")
	//viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	pb.InitJWTSecret(viper.GetString("SERVICE-ACCT-AUT.JWT_SECRET"))
}

func main() {
	err, logSync := logwriter.InitializeLoggerGRPC(
		&pb.LogGRPC, viper.GetString("SERVICE-ACCT-AUT.LOG_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	defer logSync()

	lis, err := net.Listen(
		viper.GetString("SERVICE-ACCT-AUT.NETWORK_SERVER"),
		viper.GetString("SERVICE-ACCT-AUT.PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}

	grpclog.SetLoggerV2(pb.LogGRPC)

	var (
		grpcServer = grpc.NewServer(
			grpc.ChainUnaryInterceptor(pb.InterceptotCheckCtx),
		)
		server = pb.StartServer()
	)

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
