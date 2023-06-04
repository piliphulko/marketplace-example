package aamicro

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/piliphulko/marketplace-example/internal/proto-genr/aaGrpc"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	LogGrpc *zapgrpc.Logger
)

type closeConn func()

type ConnClientMicroAa = aaGrpc.AccountAutClient

func ConnToAamicro(address string) (ConnClientMicroAa, closeConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	return aaGrpc.NewAccountAutClient(conn), func() { conn.Close() }, nil
}

type server struct {
	aaGrpc.UnimplementedAccountAutServer
	pgxPool *pgxpool.Pool
}

func StartServer() *server {
	return &server{}
}

func RegisterServer(s1 grpc.ServiceRegistrar, s2 *server) {
	aaGrpc.RegisterAccountAutServer(s1, s2)
}

func (s *server) ConnPostrgresql(postgresqlURL string) (closeConn, error) {
	pool, err := pgxpool.New(context.Background(), postgresqlURL)
	if err != nil {
		return nil, err
	}
	s.pgxPool = pool
	return func() { s.pgxPool.Close() }, nil
}
