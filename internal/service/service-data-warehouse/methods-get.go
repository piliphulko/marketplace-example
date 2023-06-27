package servicedatawarehouse

import (
	"context"
	"errors"

	"github.com/piliphulko/marketplace-example/api/basic"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrMissingMetadata = errors.New("Missing metadata")
)

func (s *server) GetAcctInfo(ctx context.Context, in *emptypb.Empty) (*basic.WarehouseInfo, error) {
	login, err := TakeLoginAndCheckJWTfromMetadataCtx(ctx)
	if err != nil {
		return nil, err
	}
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var v basic.WarehouseInfo
	if err := conn.QueryRow(ctx, `
		SELECT 
			table_warehouse_info.name_warehouse,
			table_warehouse_info.info_warehouse,
			table_warehouse_info.country,
			table_warehouse_info.city,
			table_warehouse_commission.commission_percentage
		FROM table_warehouse_info
		JOIN table_warehouse USING(id_warehouse)
		JOIN table_warehouse_commission USING (id_warehouse)
		WHERE table_warehouse.login_warehouse = $1; 
	`, login).Scan(&v.WarehouseName, &v.WarehouseNote, &v.WarehouseCountry, &v.WarehouseCity, &v.WarehouseCommission); err != nil {
		return nil, err
	}
	return &v, nil
}

func (s *server) GetInfoWallet(ctx context.Context, in *emptypb.Empty) (*basic.WarehouseWalletInfo, error) {
	login, err := TakeLoginAndCheckJWTfromMetadataCtx(ctx)
	if err != nil {
		return nil, err
	}
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var v basic.WarehouseWalletInfo
	if err := conn.QueryRow(ctx, `
		SELECT 
			table_warehouse_wallet.amount_money,
			table_warehouse_wallet.blocked_money,
			table_warehouse_commission.commission_percentage
		FROM table_warehouse_wallet
		JOIN table_warehouse USING(id_warehouse)
		JOIN table_warehouse_commission USING(id_warehouse)
		WHERE table_warehouse.login_warehouse = $1;
	`, login).Scan(&v.WalletMoneyAvailable, &v.WalletMoneyBlocked, &v.CommissionPercentage); err != nil {
		return nil, err
	}
	return &v, nil
}

func (s *server) GetArrayOrdersCustomer(ctx context.Context, in *emptypb.Empty) (*basic.ArrayOrdersCustomer, error) {
	login, err := TakeLoginAndCheckJWTfromMetadataCtx(ctx)
	if err != nil {
		return nil, err
	}
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var (
		ordersCustomer                                                 []*basic.OrdersCustomer
		orders                                                         []*basic.Order
		loginCustomer, orderUuid, nameGoods, nameVendor, nameWarehouse string
		amountGoods                                                    uint32
		priceGoods                                                     float32
	)
	rows, err := conn.Query(ctx, `
		SELECT
			table_customer.login_customer,
			table_orders.operation_uuid,
			table_goods.name_goods,
			table_vendor_info.name_vendor,
			table_warehouse_info.name_warehouse,
			table_orders.amount_goods,
			table_orders.price_goods
		FROM table_orders
		JOIN table_customer USING (id_customer)
		JOIN table_warehouse USING(id_warehouse)
		JOIN table_goods USING(id_goods)
		JOIN table_vendor_info(id_vendor)
		JOIN table_warehouse_info(id_warehouse)
		WHERE table_orders.delivery_status_order = 'confirmed order'::enum_status_order
			AND table_warehouse.login_warehouse = $1
		ORDER BY table_customer.login_customer, table_orders.operation_uuid;
	`, login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var forLoginCustomer string = ""
	for rows.Next() {
		if err := rows.Scan(&loginCustomer, &orderUuid, &nameGoods, &nameVendor, &nameWarehouse, &amountGoods, &priceGoods); err != nil {
			return nil, err
		}
		if forLoginCustomer != loginCustomer {
			ordersCustomer = append(ordersCustomer, &basic.OrdersCustomer{
				LoginCustomer: forLoginCustomer,
				Orders:        orders,
			})
			orders = nil
			orders = append(orders, &basic.Order{
				OrderUuid:     orderUuid,
				NameGoods:     nameGoods,
				NameVendor:    nameVendor,
				NameWarehouse: nameWarehouse,
				AmountGoods:   amountGoods,
				PriceGoods:    priceGoods,
			})
			forLoginCustomer = loginCustomer
		} else {
			orders = append(orders, &basic.Order{
				OrderUuid:     orderUuid,
				NameGoods:     nameGoods,
				NameVendor:    nameVendor,
				NameWarehouse: nameWarehouse,
				AmountGoods:   amountGoods,
				PriceGoods:    priceGoods,
			})
			forLoginCustomer = loginCustomer
		}
	}
	if rows.Err() != nil {
		return nil, err
	}
	if forLoginCustomer != "" {
		ordersCustomer = append(ordersCustomer, &basic.OrdersCustomer{
			LoginCustomer: forLoginCustomer,
			Orders:        orders,
		})
	}

	return &basic.ArrayOrdersCustomer{OrdersCustomer: ordersCustomer}, nil
}
