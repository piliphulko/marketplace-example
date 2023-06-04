package aamicro

import (
	"context"

	"github.com/piliphulko/marketplace-example/internal/proto-genr/basic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) AutAccount(ctx context.Context, loginPass *basic.LoginPass) (*basic.Reply, error) {
	if customerLoginPass := loginPass.GetCustomerLoginPass(); customerLoginPass != nil {
		// AUTHENTICATION CUSTOMER
		tag, err := s.pgxPool.Exec(ctx, `
		SELECT true::bool
		FROM table_customer
		WHERE login_customer = $1 AND passwort_customer = $2`,
			customerLoginPass.LoginCustomer, customerLoginPass.PasswortCustomer)
		if err != nil {
			return nil, status.New(codes.Internal, "").Err()
		}
		if tag.RowsAffected() == 0 {
			return &basic.Reply{Reply: basic.REPLY_UNAUTHORIZED}, status.New(codes.OK, "").Err()
		} else {
			return &basic.Reply{
				Reply: basic.REPLY_AUTHORIZED,
			}, status.New(codes.OK, "").Err()
		}

	} else if warehouseLoginPass := loginPass.GetWarehouseLoginPass(); warehouseLoginPass != nil {
		// AUTHENTICATION WAREHOUSE
		tag, err := s.pgxPool.Exec(ctx, `
		SELECT true::bool
		FROM table_warehouse
		WHERE login_warehouse = $1 AND passwort_warehouse = $2`,
			warehouseLoginPass.LoginWarehouse, warehouseLoginPass.PasswortWarehouse)
		if err != nil {
			return nil, status.New(codes.Internal, "").Err()
		}
		if tag.RowsAffected() == 0 {
			return &basic.Reply{Reply: basic.REPLY_UNAUTHORIZED}, status.New(codes.OK, "").Err()
		} else {
			return &basic.Reply{
				Reply: basic.REPLY_AUTHORIZED,
			}, status.New(codes.OK, "").Err()
		}

	} else if vendorLoginPass := loginPass.GetVendorLoginPass(); vendorLoginPass != nil {
		// AUTHENTICATION VENDOR
		tag, err := s.pgxPool.Exec(ctx, `
		SELECT true::bool
		FROM table_vendor
		WHERE login_vendor = $1 AND passwort_vendor = $2`,
			vendorLoginPass.LoginVendor, vendorLoginPass.PasswortVendor)
		if err != nil {
			return nil, status.New(codes.Internal, "").Err()
		}
		if tag.RowsAffected() == 0 {
			return &basic.Reply{Reply: basic.REPLY_UNAUTHORIZED}, status.New(codes.OK, "").Err()
		} else {
			return &basic.Reply{
				Reply: basic.REPLY_AUTHORIZED,
			}, status.New(codes.OK, "").Err()
		}

	} else {

		return &basic.Reply{
			Reply: basic.REPLY_UNSPECIFIED,
		}, status.New(codes.DataLoss, "").Err()

	}
}
