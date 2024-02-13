package serviceorderpay

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/api/basiccheck"
	"github.com/piliphulko/marketplace-example/internal/pkg/grpctools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) queryFunc(ctx context.Context, query, uuidV string) error {
	jwtString, err := grpctools.TakeJWTfromMetadata(ctx)
	if err != nil {
		return err
	}
	login, err := grpctools.TakeLoginAndCheckJWT(jwtString)
	if err != nil {
		return err
	}
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	var result string
	if err := conn.QueryRow(ctx, query, login, uuidV).Scan(&result); err != nil {
		return err
	}
	if result != "ok" {
		return handlerQueryError(result)
	}
	return nil
}

func (s *server) CancelOrder(ctx context.Context, in *basic.OrderUuid) (*emptypb.Empty, error) {
	if !basiccheck.NillOrNull[*basic.OrderUuid](in) {
		return &emptypb.Empty{}, apierror.ErrEmpty
	}
	const query string = `
		SELECT * FROM function_cancellation_order(
			$1::varchar,
			$2::varchar
		);`
	if err := s.queryFunc(ctx, query, in.OrderUuid); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}

func (s *server) ConfirmOrder(ctx context.Context, in *basic.OrderUuid) (*emptypb.Empty, error) {
	if !basiccheck.NillOrNull[*basic.OrderUuid](in) {
		return &emptypb.Empty{}, apierror.ErrEmpty
	}
	const query string = `
		SELECT * FROM function_confirm_order(
			$1::varchar,
			$2::varchar
		);`
	if err := s.queryFunc(ctx, query, in.OrderUuid); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}

func (s *server) CompleteOrder(ctx context.Context, in *basic.OrderUuid) (*emptypb.Empty, error) {
	if !basiccheck.NillOrNull[*basic.OrderUuid](in) {
		return &emptypb.Empty{}, apierror.ErrEmpty
	}
	const query string = `
		SELECT * FROM function_complete_order(
			$1::varchar,
			$2::varchar
		);`
	if err := s.queryFunc(ctx, query, in.OrderUuid); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}
