package donkeyhealth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type healthServer struct {
	grpc_health_v1.UnimplementedHealthServer
	pgConn *pgxpool.Pool
}

type ServiceAsFollows struct {
	Postgresql *pgxpool.Pool
}

func CreateFragmentForCheckingHealthServer(saf ServiceAsFollows) *healthServer {
	var thereIsReturning healthServer

	if saf.Postgresql != nil {
		thereIsReturning.pgConn = saf.Postgresql
	}

	return &thereIsReturning
}

func (s *healthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {

	// THERE IS CHEKING THE POSTGRESQL CONNECTION
	if err := s.pgConn.Ping(ctx); err != nil {
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
		}, nil
	} else {
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_SERVING,
		}, nil
	}
}
