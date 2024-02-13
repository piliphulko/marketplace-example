package main

import (
	"log"
	"net"

	"github.com/piliphulko/marketplace-example/internal/pkg/f16"
	pb "github.com/piliphulko/marketplace-example/internal/service/service-acct-auth"
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
	//logSync, err := logwriter.InitZapLogGRPC(
	//	&pb.LogGRPC, viper.GetString("SERVICE-ACCT-AUTH.LOG_FILE"), zapcore.ErrorLevel,
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer logSync()

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
