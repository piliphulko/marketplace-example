package microserveraccountauthentication

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/crypto/argon2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) UpdateAccount(ctx context.Context, in *basic.AccountInfoChange) (*emptypb.Empty, error) {
	var err error
	if customer := in.GetCustomerChange(); customer != nil {
		// CUSTOMER
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto errorHandling
		}
		defer conn.Release()
		var (
			bytesArgon2 []byte
			id_customer int
		)
		err = conn.QueryRow(ctx, `
			SELECT passwort_customer, id_customer
			FROM table_customer
			WHERE login_customer = $1`, customer.CustomerAutOld.LoginCustomer).Scan(&bytesArgon2, &id_customer)
		if err != nil {
			goto errorHandling
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			goto errorHandling
		}
		if !pa.CheckPass([]byte(customer.CustomerAutOld.PasswortCustomer)) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto errorHandling
		}
		defer tx.Rollback(context.TODO())
		// UPDATE
		if customer.CustomerAutNew.LoginCustomer != customer.CustomerAutOld.LoginCustomer && customer.CustomerAutNew.LoginCustomer != "" {
			_, err := tx.Exec(ctx, `
					UPDATE table_customer
					SET login_customer = $1
					WHERE id_customer = $2;`, customer.CustomerAutNew.LoginCustomer, id_customer)
			if err != nil {
				goto errorHandling
			}
		}
		if customer.CustomerAutNew.PasswortCustomer != customer.CustomerAutOld.PasswortCustomer && customer.CustomerAutNew.PasswortCustomer != "" {
			pa, err = argon2.CreareArgon2Record([]byte(customer.CustomerAutNew.PasswortCustomer))
			if err != nil {
				goto errorHandling
			}
			newPass, err := pa.Bytes()
			if err != nil {
				goto errorHandling
			}
			_, err = tx.Exec(ctx, `
					UPDATE table_customer
					SET passwort_customer = $1
					WHERE id_customer = $2;`, newPass, id_customer)
			if err != nil {
				goto errorHandling
			}
		}

		_, err = tx.Exec(ctx, `
			UPDATE table_customer_info
			SET delivery_location_country = $1::enum_country, delivery_location_city = $2
			WHERE id_customer = $3;`, customer.CustomerInfo.CustomerCountry, customer.CustomerInfo.CustomerCiry, id_customer)
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

		// WAREHOUSE
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto errorHandling
		}
		defer conn.Release()
		var (
			bytesArgon2  []byte
			id_warehouse int
		)
		err = conn.QueryRow(ctx, `
			SELECT passwort_warehouse, id_warehouse
			FROM table_warehouse
			WHERE login_warehouse = $1`, warehouse.WarehouseAutOld.LoginWarehouse).Scan(&bytesArgon2, &id_warehouse)
		if err != nil {
			goto errorHandling
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			goto errorHandling
		}
		if !pa.CheckPass([]byte(warehouse.WarehouseAutOld.PasswortWarehouse)) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto errorHandling
		}
		defer tx.Rollback(context.TODO())
		// UPDATE
		if warehouse.WarehouseAutNew.PasswortWarehouse != warehouse.WarehouseAutOld.PasswortWarehouse && warehouse.WarehouseAutNew.PasswortWarehouse != "" {
			pa, err = argon2.CreareArgon2Record([]byte(warehouse.WarehouseAutNew.PasswortWarehouse))
			if err != nil {
				goto errorHandling
			}
			newPass, err := pa.Bytes()
			if err != nil {
				goto errorHandling
			}
			_, err = tx.Exec(ctx, `
				UPDATE table_warehouse
				SET passwort_warehouse = $1
				WHERE id_warehouse = $2;`, newPass, id_warehouse)
			if err != nil {
				goto errorHandling
			}
		}
		if warehouse.WarehouseAutOld.LoginWarehouse != warehouse.WarehouseAutNew.LoginWarehouse && warehouse.WarehouseAutNew.LoginWarehouse != "" {
			_, err := tx.Exec(ctx, `
				UPDATE table_warehouse
				SET login_warehouse = $1
				WHERE id_warehouse = $2;`, warehouse.WarehouseAutNew.LoginWarehouse, id_warehouse)
			if err != nil {
				goto errorHandling
			}
		}
		var (
			SET        string = "SET "
			changeBool bool
		)
		if warehouse.WarehouseInfo.WarehouseName != "" {
			SET = fmt.Sprintf("%s name_warehouse = %s,", SET, warehouse.WarehouseInfo.WarehouseName)
			changeBool = true
		}
		if warehouse.WarehouseInfo.WarehouseNote != "" {
			SET = fmt.Sprintf("%s info_warehouse = %s,", SET, warehouse.WarehouseInfo.WarehouseNote)
			changeBool = true
		}
		if warehouse.WarehouseInfo.WarehouseCountry != "" {
			SET = fmt.Sprintf("%s country = %s::enum_country,", SET, warehouse.WarehouseInfo.WarehouseCountry)
			changeBool = true
		}
		if warehouse.WarehouseInfo.WarehouseCity != "" {
			SET = fmt.Sprintf("%s city = %s,", SET, warehouse.WarehouseInfo.WarehouseCity)
			changeBool = true
		}
		if changeBool {
			_, err = tx.Exec(ctx, ""+
				"UPDATE table_warehouse\n"+
				SET[:len(SET)-1]+"\n"+ // cut ','
				"WHERE id_warehouse = $1;", id_warehouse)
			if err != nil {
				goto errorHandling
			}
		}

		if warehouse.WarehouseInfo.WarehouseCommission != float32(0) {
			_, err = tx.Exec(ctx, `
			UPDATE table_warehouse_commission
			SET commission_percentage = $1::domain_percentage
			WHERE id_warehouse = $2;`, warehouse.WarehouseInfo.WarehouseCommission, id_warehouse)
			if err != nil {
				goto errorHandling
			}
		}
		err = tx.Commit(ctx)
		if err != nil {
			goto errorHandling
		}
		// OK
		return &emptypb.Empty{}, status.New(codes.OK, "").Err()

	} else if vendor := in.GetVendorChange(); vendor != nil {

		// VENDOR
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			goto errorHandling
		}
		defer conn.Release()
		var (
			bytesArgon2 []byte
			id_vendor   int
		)
		err = conn.QueryRow(ctx, `
			SELECT passwort_vendor, id_vendor
			FROM table_vendor
			WHERE login_vendor = $1`, vendor.VendorAutOld.LoginVendor).Scan(&bytesArgon2, &id_vendor)
		if err != nil {
			goto errorHandling
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			goto errorHandling
		}
		if !pa.CheckPass([]byte(vendor.VendorAutOld.PasswortVendor)) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		tx, err := conn.Begin(ctx)
		if err != nil {
			goto errorHandling
		}
		defer tx.Rollback(context.TODO())
		// UPDATE
		if vendor.VendorAutNew.PasswortVendor != vendor.VendorAutOld.PasswortVendor && vendor.VendorAutNew.PasswortVendor != "" {
			pa, err = argon2.CreareArgon2Record([]byte(warehouse.WarehouseAutNew.PasswortWarehouse))
			if err != nil {
				goto errorHandling
			}
			newPass, err := pa.Bytes()
			if err != nil {
				goto errorHandling
			}
			_, err = tx.Exec(ctx, `
				UPDATE table_vendor
				SET passwort_vendor = $1
				WHERE id_vendor = $2;`, newPass, id_vendor)
			if err != nil {
				goto errorHandling
			}
		}
		if vendor.VendorAutNew.LoginVendor != vendor.VendorAutOld.LoginVendor && vendor.VendorAutNew.LoginVendor != "" {
			_, err := tx.Exec(ctx, `
				UPDATE table_vendor
				SET login_vendor = $1
				WHERE id_vendor = $2;`, vendor.VendorAutNew.LoginVendor, id_vendor)
			if err != nil {
				goto errorHandling
			}
		}
		if vendor.VendorInfo.VendorName != "" {
			_, err := tx.Exec(ctx, `
				UPDATE table_vendor_info
				SET name_vendor = $1
				WHERE id_vendor = $2;`, vendor.VendorInfo.VendorName, id_vendor)
			if err != nil {
				goto errorHandling
			}
		}
		err = tx.Commit(ctx)
		if err != nil {
			goto errorHandling
		}
		// OK
		return &emptypb.Empty{}, status.New(codes.OK, "").Err()
	}
errorHandling:
	if err == pgx.ErrNoRows {
		return &emptypb.Empty{}, status.New(codes.Unauthenticated, ErrIncorrectLogin.Error()).Err()
	}
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
