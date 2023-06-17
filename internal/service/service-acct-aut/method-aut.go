package microserveraccountauthentication

import (
	"context"
	"time"

	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/crypto/argon2"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) AutAccount(ctx context.Context, loginPass *basic.LoginPass) (*basic.StringJWT, error) {
	if customerLoginPass := loginPass.GetCustomerLoginPass(); customerLoginPass != nil {
		// CUSTOMER
		// CHECK EMPTY
		if customerLoginPass.LoginCustomer == "" || &customerLoginPass.LoginCustomer == nil ||
			customerLoginPass.PasswortCustomer == "" || &customerLoginPass.PasswortCustomer == nil {
			return &basic.StringJWT{}, status.New(codes.InvalidArgument, ErrEmpty.Error()).Err()
		}
		// CHECK LENGTH PASSWORT
		if length := len(customerLoginPass.PasswortCustomer); length < 8 && length > 64 {
			return &basic.StringJWT{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
		}
		// getting password from database
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		defer conn.Release()
		var bytesArgon2 []byte
		err = conn.QueryRow(ctx, `
			SELECT passwort_customer
			FROM table_customer
			WHERE login_customer = $1`, customerLoginPass.LoginCustomer).Scan(&bytesArgon2)
		if err != nil {
			return &basic.StringJWT{}, errorHandling(err)
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		if !pa.CheckPass([]byte(customerLoginPass.PasswortCustomer)) {
			return &basic.StringJWT{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		// JWT creation
		jws, err := jwt.CreateJWS(
			jwt.Header{Alg: "SHA256", Typ: "JWT"},
			jwt.Payload{Nickname: customerLoginPass.LoginCustomer, Exp: time.Now().Add(24 * 7 * time.Hour).Unix()},
		)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		newJwt, err := jws.SignJWS()
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		return &basic.StringJWT{StringJwt: newJwt.String()}, status.New(codes.OK, "").Err()

	} else if warehouseLoginPass := loginPass.GetWarehouseLoginPass(); warehouseLoginPass != nil {
		// WAREHOUSE
		// CHECK EMPTY
		if warehouseLoginPass.LoginWarehouse == "" || &warehouseLoginPass.LoginWarehouse == nil ||
			warehouseLoginPass.PasswortWarehouse == "" || &warehouseLoginPass.PasswortWarehouse == nil {
			return &basic.StringJWT{}, status.New(codes.InvalidArgument, ErrEmpty.Error()).Err()
		}
		// CHECK LENGTH PASSWORT
		if length := len(warehouseLoginPass.PasswortWarehouse); length < 8 && length > 64 {
			return &basic.StringJWT{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
		}
		// getting password from database
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		defer conn.Release()
		var bytesArgon2 []byte
		err = conn.QueryRow(ctx, `
			SELECT passwort_warehouse
			FROM table_warehouse
			WHERE login_warehouse = $1`, warehouseLoginPass.LoginWarehouse).Scan(&bytesArgon2)
		if err != nil {
			return &basic.StringJWT{}, errorHandling(err)
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		if !pa.CheckPass([]byte(warehouseLoginPass.PasswortWarehouse)) {
			return &basic.StringJWT{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		// JWT creation
		jws, err := jwt.CreateJWS(
			jwt.Header{Alg: "SHA256", Typ: "JWT"},
			jwt.Payload{Nickname: warehouseLoginPass.LoginWarehouse, Exp: time.Now().Add(24 * 7 * time.Hour).Unix()},
		)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		newJwt, err := jws.SignJWS()
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		return &basic.StringJWT{StringJwt: newJwt.String()}, status.New(codes.OK, "").Err()

	} else if vendorLoginPass := loginPass.GetVendorLoginPass(); vendorLoginPass != nil {
		// VENDOR
		// CHECK EMPTY
		if vendorLoginPass.LoginVendor == "" || &vendorLoginPass.LoginVendor == nil ||
			vendorLoginPass.PasswortVendor == "" || &vendorLoginPass.PasswortVendor == nil {
			return &basic.StringJWT{}, status.New(codes.InvalidArgument, ErrEmpty.Error()).Err()
		}
		// CHECK LENGTH PASSWORT
		if length := len(vendorLoginPass.PasswortVendor); length < 8 && length > 64 {
			return &basic.StringJWT{}, status.New(codes.InvalidArgument, ErrPassLen.Error()).Err()
		}
		// getting password from database
		conn, err := s.AcquireConn(ctx)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		defer conn.Release()
		var bytesArgon2 []byte
		err = conn.QueryRow(ctx, `
			SELECT passwort_vendor
			FROM table_vendor
			WHERE login_vendor = $1`, vendorLoginPass.LoginVendor).Scan(&bytesArgon2)
		if err != nil {
			return &basic.StringJWT{}, errorHandling(err)
		}
		// password check
		pa, err := argon2.GetParamsArgon2(bytesArgon2)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		if !pa.CheckPass([]byte(vendorLoginPass.PasswortVendor)) {
			return &basic.StringJWT{}, status.New(codes.Unauthenticated, ErrIncorrectPass.Error()).Err()
		}
		// JWT creation
		jws, err := jwt.CreateJWS(
			jwt.Header{Alg: "SHA256", Typ: "JWT"},
			jwt.Payload{Nickname: vendorLoginPass.LoginVendor, Exp: time.Now().Add(24 * 7 * time.Hour).Unix()},
		)
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		newJwt, err := jws.SignJWS()
		if err != nil {
			LogGRPC.Error(err)
			return &basic.StringJWT{}, status.New(codes.Internal, "").Err()
		}
		return &basic.StringJWT{StringJwt: newJwt.String()}, status.New(codes.OK, "").Err()
	} else {
		LogGRPC.Info(codes.DataLoss.String())
		return &basic.StringJWT{}, status.New(codes.DataLoss, "").Err()
	}
}
