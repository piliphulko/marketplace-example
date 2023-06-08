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

func (s *server) CreateAccount(ctx context.Context, in *basic.AccountInfo) (*emptypb.Empty, error) {
	if accountInfoCustomer := in.GetCustomerInfo(); accountInfoCustomer != nil {
		// PASSWORT ENCODING
		pa, err := argon2.CreareArgon2Record([]byte(accountInfoCustomer.CustomerAut.PasswortCustomer))
		if err != nil {
			goto internalError
		}
		encodedPass, err := pa.Bytes()
		if err != nil {
			goto internalError
		}
		// DATABASE ENTY
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto internalError
		}
		defer conn.Release()
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto internalError
		}
		defer tx.Rollback(context.TODO())
		var id_customer int
		err = tx.QueryRow(ctx, `
		INSERT INTO table_customer(login_customer, passwort_customer)
		VALUES ($1, $2)
		RETURNING id_customer;`, accountInfoCustomer.CustomerAut.LoginCustomer, encodedPass).Scan(&id_customer)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				// UNIQUE ERROR
				if pgErr.Code == "23505" {
					return nil, status.New(codes.AlreadyExists, "").Err()
				}
			} else {
				goto internalError
			}
		}
		_, err = tx.Exec(ctx, `
		INSERT INTO table_customer_info (id_customer, delivery_location_country, delivery_location_city)
		VALUES ($1, $2, $3);

		INSERT INTO table_customer_wallet (id_customer) VALUES ($1);
		`, id_customer, accountInfoCustomer.CustomerCountry, accountInfoCustomer.CustomerCiry)
		if err != nil {
			goto internalError
		}
		err = tx.Commit(ctx)
		if err != nil {
			goto internalError
		}
		// OK
		return nil, status.New(codes.OK, "").Err()

	} else if accountInfoWarehouse := in.GetWarehouseInfo(); accountInfoWarehouse != nil {

		// PASSWORT ENCODING
		pa, err := argon2.CreareArgon2Record([]byte(accountInfoWarehouse.WarehouseAut.PasswortWarehouse))
		if err != nil {
			goto internalError
		}
		encodedPass, err := pa.Bytes()
		if err != nil {
			goto internalError
		}
		// DATABASE ENTY
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto internalError
		}
		defer conn.Release()
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto internalError
		}
		defer tx.Rollback(context.TODO())
		var id_warehouse int
		err = tx.QueryRow(ctx, `
		INSERT INTO table_warehouse(login_warehouse, passwort_warehouse)
		VALUES ($1, $2)
		RETURNING id_warehouse;`, accountInfoWarehouse.WarehouseAut.LoginWarehouse, encodedPass).Scan(&id_warehouse)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				// UNIQUE ERROR
				if pgErr.Code == "23505" {
					return nil, status.New(codes.AlreadyExists, "").Err()
				}
			} else {
				goto internalError
			}
		}
		_, err = tx.Exec(ctx, `
		INSERT INTO table_warehouse_wallet (id_warehouse) VALUES ($1);

		INSERT INTO table_warehouse_info (id_warehouse, name_warehouse, info_warehouse, country, city)
		VALUES ($1, $2, $3, $4::enum_country, $5);

		INSERT INTO table_warehouse_commission (id_warehouse, commission_percentage) VALUES ($1, $6::domain_percentage);
		`, id_warehouse, accountInfoWarehouse.WarehouseName, accountInfoWarehouse.WarehouseNote,
			accountInfoWarehouse.WarehouseCountry, accountInfoWarehouse.WarehouseCity, accountInfoWarehouse.WarehouseCommission)
		if err != nil {
			goto internalError
		}
		err = tx.Commit(ctx)
		if err != nil {
			goto internalError
		}
		// OK
		return nil, status.New(codes.OK, "").Err()

	} else if accountInfoVendor := in.GetVendorInfo(); accountInfoVendor != nil {

		// PASSWORT ENCODING
		pa, err := argon2.CreareArgon2Record([]byte(accountInfoVendor.VendorAut.PasswortVendor))
		if err != nil {
			goto internalError
		}
		encodedPass, err := pa.Bytes()
		if err != nil {
			goto internalError
		}
		// DATABASE ENTY
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto internalError
		}
		defer conn.Release()
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto internalError
		}
		defer tx.Rollback(context.TODO())
		var id_vendor int
		err = tx.QueryRow(ctx, `
		INSERT INTO table_vendor(login_vendor, passwort_vendor)
		VALUES ($1, $2)
		RETURNING id_vendor;`, accountInfoVendor.VendorAut.LoginVendor, encodedPass).Scan(&id_vendor)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				// UNIQUE ERROR
				if pgErr.Code == "23505" {
					return nil, status.New(codes.AlreadyExists, "").Err()
				}
			} else {
				goto internalError
			}
		}
		_, err = tx.Exec(ctx, `
		INSERT INTO table_vendor_info (id_vendor, name_vendor)
		VALUES ($1, $2);

		INSERT INTO table_vendor_wallet (id_vendor) VALUES ($1);
		`, id_vendor, accountInfoVendor.VendorName)
		if err != nil {
			goto internalError
		}
		err = tx.Commit(ctx)
		if err != nil {
			goto internalError
		}
		// OK
		return nil, status.New(codes.OK, "").Err()

	} else {
		return nil, status.New(codes.DataLoss, "").Err()
	}
	// INTERNAL ERROR
internalError:
	return nil, status.New(codes.Internal, "").Err()
}
