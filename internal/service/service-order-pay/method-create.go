package serviceorderpay

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/api/basiccheck"
	"github.com/piliphulko/marketplace-example/internal/pkg/grpctools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateOrder(ctx context.Context, in *basic.NewOrderARRAY) (*basic.OrderUuid, error) {
	arrayOrder := in.NewOrder
	if &arrayOrder == nil {
		return &basic.OrderUuid{}, apierror.ErrEmpty
	}
	jwtString, err := grpctools.TakeJWTfromMetadata(ctx)
	if err != nil {
		return &basic.OrderUuid{}, err
	}
	login, err := grpctools.TakeLoginAndCheckJWT(jwtString)
	if err != nil {
		return &basic.OrderUuid{}, err
	}
	newUuid, err := uuid.NewV4()
	if err != nil {
		return &basic.OrderUuid{}, err
	}
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return &basic.OrderUuid{}, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return &basic.OrderUuid{}, err
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
			return &basic.OrderUuid{}, apierror.ErrEmpty
		}
		var queryResult string
		if err := conn.QueryRow(ctx, query,
			newUuid, login, v.NameWarehouse, v.NameVendor, v.NameGoods, v.AmountGoods).Scan(&queryResult); err != nil {
			return &basic.OrderUuid{}, err
		}
		if queryResult != "ok" {
			return &basic.OrderUuid{}, handlerQueryError(queryResult)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return &basic.OrderUuid{}, err
	}
	return &basic.OrderUuid{OrderUuid: newUuid.String()}, status.New(codes.OK, "").Err()
}
