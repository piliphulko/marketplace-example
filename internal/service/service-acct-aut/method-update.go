package serviceacctaut

import (
	"context"
	"fmt"

	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/crypto/argon2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) UpdateAccount(ctx context.Context, in *basic.AccountInfoChange) (*emptypb.Empty, error) {

	if customer := in.GetCustomerChange(); customer != nil {
		// CUSTOMER
		if customer.CustomerAutOld == nil ||
			customer.CustomerAutOld.LoginCustomer == "" || &customer.CustomerAutOld.LoginCustomer == nil ||
			customer.CustomerAutOld.PasswortCustomer == "" || &customer.CustomerAutOld.PasswortCustomer == nil {
			return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrEmpty.Error()).Err()
		}
		// CHECK LENGTH OLD PASSWORT
		if length := len(customer.CustomerAutOld.PasswortCustomer); length < 8 && length > 64 {
			return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
		}
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
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
			return &emptypb.Empty{}, errorHandling(err)
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		if !pa.CheckPass([]byte(customer.CustomerAutOld.PasswortCustomer)) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		tx, err := conn.Begin(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		defer tx.Rollback(context.TODO())
		// UPDATE
		if customer.CustomerAutNew.LoginCustomer != customer.CustomerAutOld.LoginCustomer &&
			customer.CustomerAutNew.LoginCustomer != "" && &customer.CustomerAutNew.LoginCustomer != nil {
			_, err := tx.Exec(ctx, `
				UPDATE table_customer
				SET login_customer = $1
				WHERE id_customer = $2;`, customer.CustomerAutNew.LoginCustomer, id_customer)
			if err != nil {
				return &emptypb.Empty{}, errorHandling(err)
			}
		}
		if customer.CustomerAutNew != nil {
			if customer.CustomerAutNew.PasswortCustomer != customer.CustomerAutOld.PasswortCustomer &&
				customer.CustomerAutNew.PasswortCustomer != "" && &customer.CustomerAutNew.PasswortCustomer != nil {
				// CHECK LENGTH OLD PASSWORT
				if length := len(customer.CustomerAutNew.PasswortCustomer); length < 8 || length > 64 {
					return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
				}
				// RECORDING NEW PASSWORT
				pa, err = argon2.CreareArgon2Record([]byte(customer.CustomerAutNew.PasswortCustomer))
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
				newPass, err := pa.Bytes()
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
				_, err = tx.Exec(ctx, `
						UPDATE table_customer
						SET passwort_customer = $1
						WHERE id_customer = $2;`, newPass, id_customer)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}
		}
		if customer.CustomerInfo != nil {
			var (
				SET        string = "SET "
				changeBool bool
			)
			if customer.CustomerInfo.CustomerCountry != "" && &customer.CustomerInfo.CustomerCountry != nil {
				SET = fmt.Sprintf("%s delivery_location_country = '%s'::enum_country, ", SET, customer.CustomerInfo.CustomerCountry)
				changeBool = true
			}
			if customer.CustomerInfo.CustomerCiry != "" && &customer.CustomerInfo.CustomerCiry != nil {
				SET = fmt.Sprintf("%s delivery_location_city = '%s', ", SET, customer.CustomerInfo.CustomerCiry)
				changeBool = true
			}
			if changeBool {
				_, err = tx.Exec(ctx,
					"UPDATE table_customer_info\n"+
						SET[:len(SET)-2]+"\n"+
						"WHERE id_customer = $1;", id_customer)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}
		}
		err = tx.Commit(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		// OK
		return &emptypb.Empty{}, status.New(codes.OK, "").Err()

	} else if warehouse := in.GetWarehouseChange(); warehouse != nil {

		// WAREHOUSE
		if warehouse.WarehouseAutOld == nil ||
			warehouse.WarehouseAutOld.LoginWarehouse == "" || &warehouse.WarehouseAutOld.LoginWarehouse == nil ||
			warehouse.WarehouseAutOld.PasswortWarehouse == "" || &warehouse.WarehouseAutOld.PasswortWarehouse == nil {
			return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrEmpty.Error()).Err()
		}
		// CHECK LENGTH OLD PASSWORT
		if length := len(warehouse.WarehouseAutOld.PasswortWarehouse); length < 8 || length > 64 {
			return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
		}
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
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
			return &emptypb.Empty{}, errorHandling(err)
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		if !pa.CheckPass([]byte(warehouse.WarehouseAutOld.PasswortWarehouse)) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		tx, err := conn.Begin(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		defer tx.Rollback(context.TODO())
		// UPDATE
		if warehouse.WarehouseAutNew != nil {
			if warehouse.WarehouseAutNew.PasswortWarehouse != warehouse.WarehouseAutOld.PasswortWarehouse &&
				warehouse.WarehouseAutNew.PasswortWarehouse != "" && &warehouse.WarehouseAutNew.PasswortWarehouse != nil {
				// CHECK LENGTH NEW PASSWORT
				if length := len(warehouse.WarehouseAutNew.PasswortWarehouse); length < 8 || length > 64 {
					return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
				}
				// RECORDING NEW PASSWORT
				pa, err = argon2.CreareArgon2Record([]byte(warehouse.WarehouseAutNew.PasswortWarehouse))
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
				newPass, err := pa.Bytes()
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
				_, err = tx.Exec(ctx, `
					UPDATE table_warehouse
					SET passwort_warehouse = $1
					WHERE id_warehouse = $2;`, newPass, id_warehouse)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}

			if warehouse.WarehouseAutOld.LoginWarehouse != warehouse.WarehouseAutNew.LoginWarehouse &&
				warehouse.WarehouseAutNew.LoginWarehouse != "" && &warehouse.WarehouseAutNew.LoginWarehouse != nil {
				_, err := tx.Exec(ctx, `
					UPDATE table_warehouse
					SET login_warehouse = $1
					WHERE id_warehouse = $2;`, warehouse.WarehouseAutNew.LoginWarehouse, id_warehouse)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}
		}
		if warehouse.WarehouseInfo != nil {
			var (
				SET        string = "SET "
				changeBool bool
			)
			if warehouse.WarehouseInfo.WarehouseName != "" && &warehouse.WarehouseInfo.WarehouseName != nil {
				SET = fmt.Sprintf("%s name_warehouse = '%s', ", SET, warehouse.WarehouseInfo.WarehouseName)
				changeBool = true
			}
			if warehouse.WarehouseInfo.WarehouseNote != "" && &warehouse.WarehouseInfo.WarehouseNote != nil {
				SET = fmt.Sprintf("%s info_warehouse = '%s', ", SET, warehouse.WarehouseInfo.WarehouseNote)
				changeBool = true
			}
			if warehouse.WarehouseInfo.WarehouseCountry != "" && &warehouse.WarehouseInfo.WarehouseCountry != nil {
				SET = fmt.Sprintf("%s country = '%s'::enum_country, ", SET, warehouse.WarehouseInfo.WarehouseCountry)
				changeBool = true
			}
			if warehouse.WarehouseInfo.WarehouseCity != "" && &warehouse.WarehouseInfo.WarehouseCity != nil {
				SET = fmt.Sprintf("%s city = '%s', ", SET, warehouse.WarehouseInfo.WarehouseCity)
				changeBool = true
			}
			if changeBool {
				_, err = tx.Exec(ctx,
					"UPDATE table_warehouse\n"+
						SET[:len(SET)-2]+"\n"+ // cut ','
						"WHERE id_warehouse = $1;", id_warehouse)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}

			if warehouse.WarehouseInfo.WarehouseCommission != float32(0) && &warehouse.WarehouseInfo.WarehouseCommission != nil {
				_, err = tx.Exec(ctx, `
					UPDATE table_warehouse_commission
					SET commission_percentage = $1::domain_percentage
					WHERE id_warehouse = $2;`, warehouse.WarehouseInfo.WarehouseCommission, id_warehouse)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}
		}
		err = tx.Commit(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		// OK
		return &emptypb.Empty{}, status.New(codes.OK, "").Err()

	} else if vendor := in.GetVendorChange(); vendor != nil {

		// VENDOR
		if &vendor.VendorAutOld == nil ||
			vendor.VendorAutOld.LoginVendor == "" || &vendor.VendorAutOld.LoginVendor == nil ||
			vendor.VendorAutOld.PasswortVendor == "" || &vendor.VendorAutOld.PasswortVendor == nil {
			return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrEmpty.Error()).Err()
		}
		// CHECK LENGTH OLD PASSWORT
		if length := len(vendor.VendorAutOld.PasswortVendor); length < 8 || length > 64 {
			return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
		}
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
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
			return &emptypb.Empty{}, errorHandling(err)
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		if !pa.CheckPass([]byte(vendor.VendorAutOld.PasswortVendor)) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		tx, err := conn.Begin(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		defer tx.Rollback(context.TODO())
		// UPDATE
		if vendor.VendorAutNew != nil {
			if vendor.VendorAutNew.PasswortVendor != vendor.VendorAutOld.PasswortVendor &&
				vendor.VendorAutNew.PasswortVendor != "" && &vendor.VendorAutNew.PasswortVendor != nil {
				// CHECK LENGTH NEW PASSWORT
				if length := len(vendor.VendorAutNew.PasswortVendor); length < 8 || length > 64 {
					return &emptypb.Empty{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
				}
				// RECORDING NEW PASSWORT
				pa, err = argon2.CreareArgon2Record([]byte(warehouse.WarehouseAutNew.PasswortWarehouse))
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
				newPass, err := pa.Bytes()
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
				_, err = tx.Exec(ctx, `
					UPDATE table_vendor
					SET passwort_vendor = $1
					WHERE id_vendor = $2;`, newPass, id_vendor)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}
			if vendor.VendorAutNew.LoginVendor != vendor.VendorAutOld.LoginVendor &&
				vendor.VendorAutNew.LoginVendor != "" && &vendor.VendorAutNew.LoginVendor != nil {
				_, err := tx.Exec(ctx, `
					UPDATE table_vendor
					SET login_vendor = $1
					WHERE id_vendor = $2;`, vendor.VendorAutNew.LoginVendor, id_vendor)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}
		}
		if vendor.VendorAutNew != nil {
			if vendor.VendorInfo.VendorName != "" && &vendor.VendorInfo.VendorName != nil {
				_, err := tx.Exec(ctx, `
					UPDATE table_vendor_info
					SET name_vendor = $1
					WHERE id_vendor = $2;`, vendor.VendorInfo.VendorName, id_vendor)
				if err != nil {
					return &emptypb.Empty{}, errorHandling(err)
				}
			}
		}
		err = tx.Commit(ctx)
		if err != nil {
			return &emptypb.Empty{}, errorHandling(err)
		}
		// OK
		return &emptypb.Empty{}, status.New(codes.OK, "").Err()
	} else {
		LogGRPC.Info(codes.DataLoss.String())
		return &emptypb.Empty{}, status.New(codes.DataLoss, "").Err()
	}
}
