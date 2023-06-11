package microserveraccountauthentication

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/crypto/argon2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreateAccount(ctx context.Context, in *basic.AccountInfoChange) (*emptypb.Empty, error) {
	var err error
	if customer := in.GetCustomerChange(); customer != nil {
		// PASSWORT ENCODING
		pa, err := argon2.CreareArgon2Record([]byte(customer.CustomerAutNew.PasswortCustomer))
		if err != nil {
			goto errorHandling
		}
		encodedPass, err := pa.Bytes()
		if err != nil {
			goto errorHandling
		}
		// DATABASE ENTY
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto errorHandling
		}
		defer conn.Release()
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto errorHandling
		}
		defer tx.Rollback(context.TODO())
		var id_customer int
		err = tx.QueryRow(ctx, `
			INSERT INTO table_customer(login_customer, passwort_customer)
			VALUES ($1, $2)
			RETURNING id_customer;`, customer.CustomerAutNew.LoginCustomer, encodedPass).Scan(&id_customer)
		if err != nil {
			goto errorHandling
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO table_customer_info (id_customer, delivery_location_country, delivery_location_city)
			VALUES ($1, $2::enum_country, $3);
		`, id_customer, customer.CustomerInfo.CustomerCountry, customer.CustomerInfo.CustomerCiry)
		if err != nil {
			goto errorHandling
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO table_customer_wallet (id_customer) VALUES ($1);`, id_customer)
		if err != nil {
			goto errorHandling
		}
		err = tx.Commit(ctx)
		if err != nil {
			goto errorHandling
		}
		// OK
		return &emptypb.Empty{}, status.New(codes.OK, "").Err()

	} else if warehouse := in.GetWarehouseChange(); warehouse != nil {

		// PASSWORT ENCODING
		pa, err := argon2.CreareArgon2Record([]byte(warehouse.WarehouseAutNew.PasswortWarehouse))
		if err != nil {
			goto errorHandling
		}
		encodedPass, err := pa.Bytes()
		if err != nil {
			goto errorHandling
		}
		// DATABASE ENTY
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto errorHandling
		}
		defer conn.Release()
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto errorHandling
		}
		defer tx.Rollback(context.TODO())
		var id_warehouse int
		err = tx.QueryRow(ctx, `
			INSERT INTO table_warehouse(login_warehouse, passwort_warehouse)
			VALUES ($1, $2)
			RETURNING id_warehouse;`, warehouse.WarehouseAutNew.LoginWarehouse, encodedPass).Scan(&id_warehouse)
		if err != nil {
			goto errorHandling
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO table_warehouse_wallet (id_warehouse) VALUES ($1);`, id_warehouse)
		if err != nil {
			goto errorHandling
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO table_warehouse_info (id_warehouse, name_warehouse, info_warehouse, country, city)
			VALUES ($1, $2, $3, $4::enum_country, $5);`,
			id_warehouse, warehouse.WarehouseInfo.WarehouseName, warehouse.WarehouseInfo.WarehouseNote,
			warehouse.WarehouseInfo.WarehouseCountry, warehouse.WarehouseInfo.WarehouseCity)
		if err != nil {
			goto errorHandling
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO table_warehouse_commission (id_warehouse, commission_percentage) VALUES ($1, $2::domain_percentage);
		`, id_warehouse, warehouse.WarehouseInfo.WarehouseCommission)
		if err != nil {
			goto errorHandling
		}
		err = tx.Commit(ctx)
		if err != nil {
			goto errorHandling
		}
		// OK
		return nil, status.New(codes.OK, "").Err()

	} else if vendor := in.GetVendorChange(); vendor != nil {

		// PASSWORT ENCODING
		pa, err := argon2.CreareArgon2Record([]byte(vendor.VendorAutNew.PasswortVendor))
		if err != nil {
			goto errorHandling
		}
		encodedPass, err := pa.Bytes()
		if err != nil {
			goto errorHandling
		}
		// DATABASE ENTY
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto errorHandling
		}
		defer conn.Release()
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto errorHandling
		}
		defer tx.Rollback(context.TODO())
		var id_vendor int
		err = tx.QueryRow(ctx, `
			INSERT INTO table_vendor(login_vendor, passwort_vendor)
			VALUES ($1, $2)
			RETURNING id_vendor;`, vendor.VendorAutNew.LoginVendor, encodedPass).Scan(&id_vendor)
		if err != nil {
			goto errorHandling
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO table_vendor_info (id_vendor, name_vendor)
			VALUES ($1, $2);`, id_vendor, vendor.VendorInfo.VendorName)
		if err != nil {
			goto errorHandling
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO table_vendor_wallet (id_vendor) VALUES ($1);
		`, id_vendor)
		if err != nil {
			goto errorHandling
		}
		err = tx.Commit(ctx)
		if err != nil {
			goto errorHandling
		}
		// OK
		return &emptypb.Empty{}, status.New(codes.OK, "").Err()

	} else {
		return &emptypb.Empty{}, status.New(codes.DataLoss, "").Err()
	}
errorHandling:
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// UNIQUE ERROR
		if pgErr.Code == "23505" {
			return &emptypb.Empty{}, status.New(codes.AlreadyExists, "").Err()
			// INCORRECT COUNTRY
		} else if pgErr.Code == "22P02" {
			return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrIncorrectCountry.Error()).Err()
		}
	}
	return &emptypb.Empty{}, status.New(codes.Internal, "").Err()
}
