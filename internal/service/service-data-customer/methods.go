package servicedatacustomer

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/grpctools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) selectOrders(ctx context.Context, login, orderStatus string) (*basic.Orders, error) {
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return &basic.Orders{}, err
	}
	defer conn.Release()

	if orderStatus == "completed order" {
		const query = `
		SELECT 
			operation_uuid::varchar,
			CONCAT(delivery_location_country, ', ', delivery_location_city) AS delivery_location,
			name_goods,
			type_goods,
			name_vendor,
			name_warehouse,
			amount_goods,
			price_goods,
			date_order_finish
		FROM view_orders_all
		WHERE login_customer = $1 AND delivery_status_order = $2::enum_status_order;`

		rows, err := conn.Query(ctx, query, login, orderStatus)
		if err != nil {
			return &basic.Orders{}, err
		}
		defer rows.Close()
		var orders []*basic.Order

		for rows.Next() {
			var o basic.Order
			err := rows.Scan(&o.OrderUuid, &o.Location, &o.NameGoods, &o.TypeGoods, &o.NameVendor, &o.NameWarehouse, &o.AmountGoods, &o.PriceGoods, &o.Date)
			if err != nil {
				return &basic.Orders{}, err
			}
			orders = append(orders, &o)
		}
		if err := rows.Err(); err != nil {
			return &basic.Orders{}, err
		}
		return &basic.Orders{Orders: orders}, status.New(codes.OK, "").Err()

	} else {
		const query = `
		SELECT 
			operation_uuid::varchar,
			CONCAT(delivery_location_country, ', ', delivery_location_city) AS delivery_location,
			name_goods,
			type_goods,
			name_vendor,
			name_warehouse,
			amount_goods,
			price_goods
		FROM view_orders_all
		WHERE login_customer = $1 AND delivery_status_order = $2::enum_status_order;`

		rows, err := conn.Query(ctx, query, login, orderStatus)
		if err != nil {
			return &basic.Orders{}, err
		}
		defer rows.Close()
		var orders []*basic.Order

		for rows.Next() {
			var o basic.Order
			err := rows.Scan(&o.OrderUuid, &o.Location, &o.NameGoods, &o.TypeGoods, &o.NameVendor, &o.NameWarehouse, &o.AmountGoods, &o.PriceGoods)
			if err != nil {
				return &basic.Orders{}, err
			}
			orders = append(orders, &o)
		}
		if err := rows.Err(); err != nil {
			return &basic.Orders{}, err
		}
		return &basic.Orders{Orders: orders}, status.New(codes.OK, "").Err()
	}

}

func (s *server) GetCustomerOrders(ctx context.Context, in *basic.OrderStatus) (*basic.Orders, error) {
	if &in == nil {
		return &basic.Orders{}, apierror.ErrEmpty
	}

	token, err := grpctools.TakeJWTfromMetadata(ctx)
	if err != nil {
		return &basic.Orders{}, err
	}

	login, err := grpctools.TakeLoginAndCheckJWT(token)
	if err != nil {
		return &basic.Orders{}, err
	}
	switch in.OrderStatus {

	case basic.OrderStatusEnum_UNCONFIRMED:

		return s.selectOrders(ctx, login, "unconfirmed order")

	case basic.OrderStatusEnum_CONFIRNED:

		return s.selectOrders(ctx, login, "confirmed order")

	case basic.OrderStatusEnum_COMPLETED:

		return s.selectOrders(ctx, login, "completed order")

	default:
		return &basic.Orders{}, apierror.ErrEmpty
	}
}

func (s *server) GetCustomerInfo(ctx context.Context, in *emptypb.Empty) (*basic.CustomerInfo, error) {
	token, err := grpctools.TakeJWTfromMetadata(ctx)
	if err != nil {
		return &basic.CustomerInfo{}, err
	}
	login, err := grpctools.TakeLoginAndCheckJWT(token)
	if err != nil {
		return &basic.CustomerInfo{}, err
	}
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return &basic.CustomerInfo{}, err
	}
	var ci basic.CustomerInfo
	const query = `
		SELECT table_customer_info.delivery_location_country, table_customer_info.delivery_location_city
		FROM table_customer_info
		JOIN table_customer USING (id_customer)
		WHERE table_customer.login_customer = $1;`

	if err := conn.QueryRow(ctx, query, login).Scan(&ci.CustomerCountry, &ci.CustomerCity); err != nil {
		return &basic.CustomerInfo{}, err
	}
	return &basic.CustomerInfo{
		CustomerCountry: ci.CustomerCountry,
		CustomerCity:    ci.CustomerCity,
	}, status.New(codes.OK, "").Err()
}

func (s *server) GetWalletInfo(ctx context.Context, in *emptypb.Empty) (*basic.WalletInfo, error) {
	token, err := grpctools.TakeJWTfromMetadata(ctx)
	if err != nil {
		return &basic.WalletInfo{}, err
	}
	login, err := grpctools.TakeLoginAndCheckJWT(token)
	if err != nil {
		return &basic.WalletInfo{}, err
	}
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return &basic.WalletInfo{}, err
	}
	var wi basic.WalletInfo
	const query = `
		SELECT table_customer_wallet.amount_money, table_customer_wallet.blocked_money
		FROM table_customer_wallet
		JOIN table_customer USING (id_customer)
		WHERE table_customer.login_customer = $1;`

	if err := conn.QueryRow(ctx, query, login).Scan(&wi.AmountMoney, &wi.BlockedMoney); err != nil {
		return &basic.WalletInfo{}, err
	}
	return &basic.WalletInfo{
		AmountMoney:  wi.AmountMoney,
		BlockedMoney: wi.BlockedMoney,
	}, status.New(codes.OK, "").Err()
}
