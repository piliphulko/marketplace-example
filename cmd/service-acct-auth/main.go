package main

import (
	"log"
	"net"

	"github.com/piliphulko/marketplace-example/internal/pkg/donkeyhealth"
	"github.com/piliphulko/marketplace-example/internal/pkg/donkeypacking"
	"github.com/piliphulko/marketplace-example/internal/pkg/f16"
	pb "github.com/piliphulko/marketplace-example/internal/service/service-acct-auth"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct {
	grpc_health_v1.UnimplementedHealthServer
}

func init() {
	//viper.SetConfigFile("config.yaml")
	viper.SetConfigFile("/usr/local/bin/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	pb.InitJWTSecret(viper.GetString("SERVICE-ACCT-AUTH.JWT_SECRET"))
}

func main() {

	lis, err := net.Listen(
		viper.GetString("SERVICE-ACCT-AUTH.NETWORK_SERVER"),
		viper.GetString("SERVICE-ACCT-AUTH.PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}

	//grpclog.SetLoggerV2(pb.LogGRPC)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			f16.InterceptorCheckCtx,
			f16.IntrceptorHandlerErrors,
		),
	)

	pgxPool, closePgxPool, err := donkeypacking.GetConnPostrgresql(viper.GetString("POSTGRESQL.DATABASE_URL"))

	defer closePgxPool()
	if err != nil {
		log.Fatal(err)
	}

	server := pb.StartServer()

	server.InsertPostgresql(pgxPool)

	healthFragmentServer := donkeyhealth.CreateFragmentForCheckingHealthServer(
		donkeyhealth.ServiceAsFollows{
			Postgresql: pgxPool,
		},
	)

	grpc_health_v1.RegisterHealthServer(grpcServer, healthFragmentServer)

	pb.RegisterServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
