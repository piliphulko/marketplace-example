package microserveraccountauthentication

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"github.com/piliphulko/marketplace-example/internal/service/microserver-account-authentication/core"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
)

var (
	LogGRPC *zapgrpc.Logger
)

var (
	ErrIncorrectPass    = errors.New("Incorrect password")
	ErrIncorrectLogin   = errors.New("Incorrect login")
	ErrIncorrectCountry = errors.New("Incorrect country")
)

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
	pool, err := pgxpool.New(context.Background(), postgresqlURL)
	if err != nil {
		return nil, err
	}
	s.pgxPool = pool
	return func() { s.pgxPool.Close() }, nil
}

func (s *server) AcquireConn(ctx context.Context) (*pgxpool.Conn, error) {
	return s.pgxPool.Acquire(ctx)
}

func InitJWTSecret(secret string) {
	jwt.InsertSecretForSignJWS(secret)
}
