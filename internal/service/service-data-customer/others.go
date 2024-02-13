package servicedatacustomer

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"github.com/piliphulko/marketplace-example/internal/service/service-data-customer/core"
	"google.golang.org/grpc"
)

type closeConn func()

type server struct {
	core.UnimplementedDataCustomerServer
	pgxPool *pgxpool.Pool
}

func StartServer() *server {
	return &server{}
}

func RegisterServer(s1 grpc.ServiceRegistrar, s2 *server) {
	core.RegisterDataCustomerServer(s1, s2)
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
