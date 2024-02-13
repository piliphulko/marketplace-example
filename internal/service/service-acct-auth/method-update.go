package serviceacctauth

import (
	"context"
	"fmt"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/crypto/argon2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) ChangeAccount(ctx context.Context, in *basic.CustomerChange) (*emptypb.Empty, error) {
	conn, err := s.pgxPool.Acquire(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer conn.Release()
	var (
		bytesArgon2 []byte
		id_customer int
	)
	err = conn.QueryRow(ctx, `
			SELECT passwort_customer, id_customer
			FROM table_customer
			WHERE login_customer = $1`, in.CustomerAutOld.LoginCustomer).Scan(&bytesArgon2, &id_customer)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	// password check
	pa, err := argon2.GetParamsArgon2(bytesArgon2)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if !pa.CheckPass([]byte(in.CustomerAutOld.PasswordCustomer)) {
		return &emptypb.Empty{}, apierror.ErrIncorrectPass
	}
	if in.CustomerAutNew.PrimaryValidation() != nil || in.CustomerInfo.PrimaryValidation() != nil {
		return &emptypb.Empty{}, apierror.ErrNotDataUpdate
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer tx.Rollback(context.TODO())
	// UPDATE
	if in.CustomerAutNew != nil {
		if in.CustomerAutNew.LoginCustomer != in.CustomerAutOld.LoginCustomer &&
			in.CustomerAutNew.LoginCustomer != "" && &in.CustomerAutNew.LoginCustomer != nil {
			_, err := tx.Exec(ctx, `
				UPDATE table_customer
				SET login_customer = $1
				WHERE id_customer = $2;`, in.CustomerAutNew.LoginCustomer, id_customer)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}
		if in.CustomerAutNew.PasswordCustomer != in.CustomerAutOld.PasswordCustomer ||
			in.CustomerAutNew.PasswordCustomer != "" && &in.CustomerAutNew.PasswordCustomer != nil {
			// CHECK LENGTH OLD PASSWORT
			if !basic.PasswotdRule(in.CustomerAutNew.PasswordCustomer) {
				return &emptypb.Empty{}, apierror.ErrPassLen
			}
			// RECORDING NEW PASSWORT
			pa, err = argon2.CreareArgon2Record([]byte(in.CustomerAutNew.PasswordCustomer))
			if err != nil {
				return &emptypb.Empty{}, err
			}
			newPass, err := pa.Bytes()
			if err != nil {
				return &emptypb.Empty{}, err
			}
			_, err = tx.Exec(ctx, `
						UPDATE table_customer
						SET passwort_customer = $1
						WHERE id_customer = $2;`, newPass, id_customer)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}
	}
	if in.CustomerInfo != nil {
		var (
			SET        string = "SET "
			changeBool bool
		)
		if in.CustomerInfo.CustomerCountry != "" {
			SET = fmt.Sprintf("%s delivery_location_country = '%s'::enum_country, ", SET, in.CustomerInfo.CustomerCountry)
			changeBool = true
		}
		if in.CustomerInfo.CustomerCity != "" {
			SET = fmt.Sprintf("%s delivery_location_city = '%s', ", SET, in.CustomerInfo.CustomerCity)
			changeBool = true
		}
		if changeBool {
			_, err = tx.Exec(ctx,
				"UPDATE table_customer_info\n"+
					SET[:len(SET)-2]+"\n"+
					"WHERE id_customer = $1;", id_customer)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// OK
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}

func (s *server) ChangeAccountWarehouse(ctx context.Context, in *basic.WarehouseChange) (*emptypb.Empty, error) {

	conn, err := s.pgxPool.Acquire(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer conn.Release()
	var (
		bytesArgon2  []byte
		id_warehouse int
	)
	err = conn.QueryRow(ctx, `
			SELECT passwort_warehouse, id_warehouse
			FROM table_warehouse
			WHERE login_warehouse = $1`, in.WarehouseAutOld.LoginWarehouse).Scan(&bytesArgon2, &id_warehouse)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// password check
	pa, err := argon2.GetParamsArgon2(bytesArgon2)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if !pa.CheckPass([]byte(in.WarehouseAutOld.PasswordWarehouse)) {
		return &emptypb.Empty{}, apierror.ErrIncorrectPass
	}
	if in.WarehouseAutNew.PrimaryValidation() != nil || in.WarehouseInfo.PrimaryValidation() != nil {
		return &emptypb.Empty{}, apierror.ErrNotDataUpdate
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer tx.Rollback(context.TODO())
	// UPDATE
	if in.WarehouseAutNew != nil {
		if in.WarehouseAutNew.PasswordWarehouse != in.WarehouseAutOld.PasswordWarehouse &&
			in.WarehouseAutNew.PasswordWarehouse != "" && &in.WarehouseAutNew.PasswordWarehouse != nil {
			// CHECK LENGTH NEW PASSWORT
			if !basic.PasswotdRule(in.WarehouseAutNew.PasswordWarehouse) {
				return &emptypb.Empty{}, apierror.ErrPassLen
			}
			// RECORDING NEW PASSWORT
			pa, err = argon2.CreareArgon2Record([]byte(in.WarehouseAutNew.PasswordWarehouse))
			if err != nil {
				return &emptypb.Empty{}, err
			}
			newPass, err := pa.Bytes()
			if err != nil {
				return &emptypb.Empty{}, err
			}
			_, err = tx.Exec(ctx, `
					UPDATE table_warehouse
					SET passwort_warehouse = $1
					WHERE id_warehouse = $2;`, newPass, id_warehouse)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}

		if in.WarehouseAutOld.LoginWarehouse != in.WarehouseAutNew.LoginWarehouse &&
			in.WarehouseAutNew.LoginWarehouse != "" && &in.WarehouseAutNew.LoginWarehouse != nil {
			_, err := tx.Exec(ctx, `
					UPDATE table_warehouse
					SET login_warehouse = $1
					WHERE id_warehouse = $2;`, in.WarehouseAutNew.LoginWarehouse, id_warehouse)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}
	}
	if in.WarehouseInfo != nil {
		var (
			SET        string = "SET "
			changeBool bool
		)
		if in.WarehouseInfo.WarehouseName != "" && &in.WarehouseInfo.WarehouseName != nil {
			SET = fmt.Sprintf("%s name_warehouse = '%s', ", SET, in.WarehouseInfo.WarehouseName)
			changeBool = true
		}
		if in.WarehouseInfo.WarehouseNote != "" && &in.WarehouseInfo.WarehouseNote != nil {
			SET = fmt.Sprintf("%s info_warehouse = '%s', ", SET, in.WarehouseInfo.WarehouseNote)
			changeBool = true
		}
		if in.WarehouseInfo.WarehouseCountry != "" && &in.WarehouseInfo.WarehouseCountry != nil {
			SET = fmt.Sprintf("%s country = '%s'::enum_country, ", SET, in.WarehouseInfo.WarehouseCountry)
			changeBool = true
		}
		if in.WarehouseInfo.WarehouseCity != "" && &in.WarehouseInfo.WarehouseCity != nil {
			SET = fmt.Sprintf("%s city = '%s', ", SET, in.WarehouseInfo.WarehouseCity)
			changeBool = true
		}
		if changeBool {
			_, err = tx.Exec(ctx,
				"UPDATE table_warehouse\n"+
					SET[:len(SET)-2]+"\n"+ // cut ','
					"WHERE id_warehouse = $1;", id_warehouse)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}

		if in.WarehouseInfo.WarehouseCommission != float32(0) && &in.WarehouseInfo.WarehouseCommission != nil {
			_, err = tx.Exec(ctx, `
					UPDATE table_warehouse_commission
					SET commission_percentage = $1::domain_percentage
					WHERE id_warehouse = $2;`, in.WarehouseInfo.WarehouseCommission, id_warehouse)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	// OK
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}

func (s *server) ChangeAccountVendor(ctx context.Context, in *basic.VendorChange) (*emptypb.Empty, error) {

	conn, err := s.pgxPool.Acquire(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer conn.Release()
	var (
		bytesArgon2 []byte
		id_vendor   int
	)
	err = conn.QueryRow(ctx, `
	SELECT passwort_vendor, id_vendor
	FROM table_vendor
	WHERE login_vendor = $1`, in.VendorAutOld.LoginVendor).Scan(&bytesArgon2, &id_vendor)
	if err != nil {
		return &emptypb.Empty{}, handlingErrSql(err)
	}
	// password check
	pa, err := argon2.GetParamsArgon2(bytesArgon2)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if !pa.CheckPass([]byte(in.VendorAutOld.PasswordVendor)) {
		return &emptypb.Empty{}, apierror.ErrIncorrectPass
	}
	if in.VendorAutNew.PrimaryValidation() != nil || in.VendorInfo.PrimaryValidation() != nil {
		return &emptypb.Empty{}, apierror.ErrNotDataUpdate
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer tx.Rollback(context.TODO())
	// UPDATE
	if in.VendorAutNew != nil {
		if in.VendorAutNew.PasswordVendor != in.VendorAutOld.PasswordVendor &&
			in.VendorAutNew.PasswordVendor != "" && &in.VendorAutNew.PasswordVendor != nil {
			// CHECK LENGTH NEW PASSWORT
			if !basic.PasswotdRule(in.VendorAutNew.PasswordVendor) {
				return &emptypb.Empty{}, apierror.ErrPassLen
			}
			// RECORDING NEW PASSWORT
			pa, err = argon2.CreareArgon2Record([]byte(in.VendorAutNew.PasswordVendor))
			if err != nil {
				return &emptypb.Empty{}, err
			}
			newPass, err := pa.Bytes()
			if err != nil {
				return &emptypb.Empty{}, err
			}
			_, err = tx.Exec(ctx, `
			UPDATE table_vendor
			SET passwort_vendor = $1
			WHERE id_vendor = $2;`, newPass, id_vendor)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}
		if in.VendorAutNew.LoginVendor != in.VendorAutOld.LoginVendor &&
			in.VendorAutNew.LoginVendor != "" && &in.VendorAutNew.LoginVendor != nil {
			_, err := tx.Exec(ctx, `
			UPDATE table_vendor
			SET login_vendor = $1
			WHERE id_vendor = $2;`, in.VendorAutNew.LoginVendor, id_vendor)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}
	}
	if in.VendorAutNew != nil {
		if in.VendorInfo.VendorName != "" && &in.VendorInfo.VendorName != nil {
			_, err := tx.Exec(ctx, `
			UPDATE table_vendor_info
			SET name_vendor = $1
			WHERE id_vendor = $2;`, in.VendorInfo.VendorName, id_vendor)
			if err != nil {
				return &emptypb.Empty{}, handlingErrSql(err)
			}
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}
