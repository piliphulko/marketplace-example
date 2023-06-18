package serviceacctaut

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"github.com/piliphulko/marketplace-example/internal/service/service-acct-aut/core"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	LogGRPC *zapgrpc.Logger
)

var (
	ErrIncorrectPass    = errors.New("Incorrect password")
	ErrIncorrectLogin   = errors.New("Incorrect login")
	ErrIncorrectCountry = errors.New("Incorrect country")
	ErrEmpty            = errors.New("Empty value passed")
	ErrPassLen          = errors.New("password is not in the allowed number of characters (8-64)")
)

func errorHandling(err error) error {
	var pgErr *pgconn.PgError
	if err == pgx.ErrNoRows {
		return status.New(codes.Unauthenticated, ErrIncorrectLogin.Error()).Err()
	}
	LogGRPC.Error(err)
	if errors.As(err, &pgErr) {
		// UNIQUE ERROR
		if pgErr.Code == "23505" {
			return status.New(codes.AlreadyExists, "").Err()
			// INCORRECT COUNTRY
		} else if pgErr.Code == "22P02" {
			return status.New(codes.InvalidArgument, ErrIncorrectCountry.Error()).Err()
		}
	}
	return status.New(codes.Internal, "").Err()
}

type closeConn func()

type server struct {
	core.UnimplementedAccountAutServer
	pgxPool *pgxpool.Pool
}

func StartServer() *server {
	return &server{}
}

func RegisterServer(s1 grpc.ServiceRegistrar, s2 *server) {
	core.RegisterAccountAutServer(s1, s2)
}

func (s *server) ConnPostrgresql(postgresqlURL string) (closeConn, error) {
	var i int = 1
	for {
		fmt.Printf("|POSTGRESQL|:connection attempt: %d\n", i)
		pool, err := pgxpool.New(context.Background(), postgresqlURL)
		if err != nil && i > 4 {
			return nil, err
		} else if err == nil {
			fmt.Println("|POSTGRESQL|:connection completed successfully")
			s.pgxPool = pool
			return func() { s.pgxPool.Close() }, nil
		}
		time.Sleep(time.Duration(i^2*250) * time.Microsecond)
		i++
	}
}

func (s *server) AcquireConn(ctx context.Context) (*pgxpool.Conn, error) {
	return s.pgxPool.Acquire(ctx)
}

func InitJWTSecret(secret string) {
	jwt.InsertSecretForSignJWS(secret)
}
