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

func (s *server) AddToOrder(ctx context.Context, in *basic.AddToOrderARRAY) (*emptypb.Empty, error) {
	arrayOrder := in.NewOrder
	if &arrayOrder == nil {
		return &emptypb.Empty{}, apierror.ErrEmpty
	}
	jwtString, err := grpctools.TakeJWTfromMetadata(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	login, err := grpctools.TakeLoginAndCheckJWT(jwtString)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer tx.Rollback(context.TODO())
	const query string = `
		SELECT * FROM function_create_order(
			$1::varchar,
			$2::varchar,
			$3::varchar,
			$4::varchar,
			$5::varchar,
			$6::int
		);`
	for _, v := range arrayOrder {
		if !basiccheck.NillOrNull[*basic.NewOrder](v) {
			return &emptypb.Empty{}, apierror.ErrEmpty
		}
		var queryResult string
		if err := conn.QueryRow(ctx, query,
			in.OrderUuid, login, v.NameWarehouse, v.NameVendor, v.NameGoods, v.AmountGoods).Scan(&queryResult); err != nil {
			return &emptypb.Empty{}, err
		}
		if queryResult != "ok" {
			return &emptypb.Empty{}, handlerQueryError(queryResult)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}
