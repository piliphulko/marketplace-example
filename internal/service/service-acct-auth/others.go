package serviceacctauth

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"github.com/piliphulko/marketplace-example/internal/service/service-acct-auth/core"
	"google.golang.org/grpc"
)

var (
	ErrIncorrectPass    = errors.New("Incorrect password")
	ErrIncorrectLogin   = errors.New("Incorrect login")
	ErrIncorrectCountry = errors.New("Incorrect country")
	ErrEmpty            = errors.New("Empty value passed")
	ErrPassLen          = errors.New("Password is not in the allowed number of characters (8-64)")
	ErrLoginBusy        = errors.New("Login busy")
)
var JwtClaims = struct {
	Alg string
	Typ string
	Exp int64
}{
	Alg: "SHA256",
	Typ: "JWT",
	Exp: time.Now().Add(24 * 7 * time.Hour).Unix(),
}

func handlingErrSql(err error) error {
	var pgErr *pgconn.PgError
	if err == pgx.ErrNoRows {
		return apierror.ErrIncorrectLogin
	}
	if errors.As(err, &pgErr) {
		// UNIQUE ERROR
		if pgErr.Code == "23505" {
			return apierror.ErrLoginBusy
			// INCORRECT COUNTRY
		} else if pgErr.Code == "22P02" {
			return apierror.ErrIncorrectCountry
		}
	}
	return err
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

func (s *server) InsertPostgresql(pgxPool *pgxpool.Pool) { s.pgxPool = pgxPool }

func (s *server) AcquireConn(ctx context.Context) (*pgxpool.Conn, error) {
	return s.pgxPool.Acquire(ctx)
}

func InitJWTSecret(secret string) {
	jwt.InsertSecretForSignJWS(secret)
}
