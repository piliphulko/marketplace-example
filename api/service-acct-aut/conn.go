package serviceacctauth

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/service/service-acct-auth/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type closeConn func()

type AccountAutClient interface {
	core.AccountAutClient
}

type accountAutClient struct {
	core.AccountAutClient
}

func ConnToServiceAccountAuthentication(address string) (AccountAutClient, closeConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return core.NewAccountAutClient(conn), func() { conn.Close() }, nil
}

func (c *accountAutClient) CreateAccount(ctx context.Context, in *basic.CustomerNew, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.CreateAccount(ctx, in, opts...)
}

func (c *accountAutClient) CreateAccountWarehouse(ctx context.Context, in *basic.WarehouseNew, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.CreateAccountWarehouse(ctx, in, opts...)
}

func (c *accountAutClient) CreateAccountVendor(ctx context.Context, in *basic.VendorNew, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.CreateAccountVendor(ctx, in, opts...)
}

func (c *accountAutClient) AutAccount(ctx context.Context, in *basic.CustomerAut, opts ...grpc.CallOption) (*basic.StringJWT, error) {
	return c.AutAccount(ctx, in, opts...)
}

func (c *accountAutClient) AutAccountWarehouse(ctx context.Context, in *basic.WarehouseAut, opts ...grpc.CallOption) (*basic.StringJWT, error) {
	return c.AutAccountWarehouse(ctx, in, opts...)
}

func (c *accountAutClient) AutAccountVendor(ctx context.Context, in *basic.VendorAut, opts ...grpc.CallOption) (*basic.StringJWT, error) {
	return c.AutAccountVendor(ctx, in, opts...)
}

func (c *accountAutClient) ChangeAccount(ctx context.Context, in *basic.CustomerChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.ChangeAccount(ctx, in, opts...)
}

func (c *accountAutClient) ChangeAccountWarehouse(ctx context.Context, in *basic.WarehouseChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.ChangeAccountWarehouse(ctx, in, opts...)
}

func (c *accountAutClient) ChangeAccountVendor(ctx context.Context, in *basic.VendorChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.ChangeAccountVendor(ctx, in, opts...)
}

/*
type ChoiseAccount interface {
	basic.CustomerChange | basic.VendorChange | basic.WarehouseChange
}

func OneofAccount[T ChoiseAccount](v T) *basic.AccountInfoChange {
	switch kind := any(v).(type) {
	case basic.CustomerChange:
		return &basic.AccountInfoChange{
			AccountInfo: &basic.AccountInfoChange_CustomerChange{
				CustomerChange: &kind,
			}}
	case basic.VendorChange:
		return &basic.AccountInfoChange{
			AccountInfo: &basic.AccountInfoChange_VendorChange{
				VendorChange: &kind,
			}}
	case basic.WarehouseChange:
		return &basic.AccountInfoChange{
			AccountInfo: &basic.AccountInfoChange_WarehouseChange{
				WarehouseChange: &kind,
			}}
	}
	return nil
}

type ChoiseLoginPass interface {
	basic.CustomerAut | basic.VendorAut | basic.WarehouseAut
}

func OneofLoginPass[T ChoiseLoginPass](v T) *basic.LoginPass {
	switch kind := any(v).(type) {
	case basic.CustomerAut:
		return &basic.LoginPass{
			AccountChoice: &basic.LoginPass_CustomerLoginPass{
				CustomerLoginPass: &kind,
			},
		}
	case basic.VendorAut:
		return &basic.LoginPass{
			AccountChoice: &basic.LoginPass_VendorLoginPass{
				VendorLoginPass: &kind,
			},
		}
	case basic.WarehouseAut:
		return &basic.LoginPass{
			AccountChoice: &basic.LoginPass_WarehouseLoginPass{
				WarehouseLoginPass: &kind,
			},
		}
	}
	return nil
}

// ConnToMicroserverAccountAuthentication getting a connection to a service
func ConnToServiceAccountAuthentication(address string) (AccountAuthClient, closeConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return core.NewAccountAutClient(conn), func() { conn.Close() }, nil
}

// AutAccount account authentication and getting a JWT token
// user type should be selected, for easier operation, you can use the OneofLoginPass function
// possible errors: ErrEmpty, ErrPassLen, ErrIncorrectPass, ErrIncorrectLogin
/*
	pbClient "github.com/piliphulko/marketplace-example/api/service-accth-aut"

	jwtString, err := conn.AutAccount(ctx, pbClient.OneofLoginPass(basic.CustomerAut{
		LoginCustomer:    "newlogin",
		PasswortCustomer: "123456ab",
	}))
*/
/*
func (aa *accountAuthClient) AutAccount(ctx context.Context, in *basic.LoginPass, opts ...grpc.CallOption) (*basic.StringJWT, error) {
	return aa.AutAccount(ctx, in, opts...)
}

// CheckJWT checks the token, if the token is invalid then returns an error
// possible errors: ErrTokenFake, ErrTokenExpired
func (aa *accountAuthClient) CheckJWT(ctx context.Context, in *basic.StringJWT, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return aa.CheckJWT(ctx, in, opts...)
}

// CreateAccount creates an account, all fields must be filled
// for convenient data filling, you can use the OneofAccount ​​function
// country name must be in uppercase
// possible errors: ErrEmpty, ErrPassLen, ErrIncorrectCountry, ErrLoginBusy
/*
	pbClient "github.com/piliphulko/marketplace-example/api/service-accth-aut"

	_, err = conn.CreateAccount(ctx, pbClient.OneofAccount(basic.CustomerChange{
		CustomerAutNew: &basic.CustomerAut{
			LoginCustomer:    "newlogin",
			PasswortCustomer: "123456ab",
		},
		CustomerInfo: &basic.CustomerInfo{
			CustomerCountry: "BELARUS",
			CustomerCity:    "MINSK",
		},
	}))
*/
/*
func (aa *accountAuthClient) CreateAccount(ctx context.Context, in *basic.AccountInfoChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return aa.CreateAccount(ctx, in, opts...)
}
*/
// ChangeAccount changes account information
// for convenient data filling, you can use the OneofAccount ​​function
// country name must be in uppercase
// changes only filled fields
// possible errors: ErrEmpty, ErrPassLen, ErrIncorrectCountry, ErrLoginBusy, ErrIncorrectLogin, ErrIncorrectPass
/*
	pbClient "github.com/piliphulko/marketplace-example/api/service-acct-auth"

	_, err = conn.ChangeAccount(ctx, pbClient.OneofAccount(basic.CustomerChange{
		CustomerAutOld: &basic.CustomerAut{ // to change, you need to confirm the data
			LoginCustomer:    "newlogin",
			PasswortCustomer: "123456ab",
		},
		CustomerAutNew: &basic.CustomerAut{
			LoginCustomer:    "newlogin2", // this request will only change the login
		},
		CustomerInfo: &basic.CustomerInfo{
			CustomerCountry: "BELARUS", // passing an active value will not raise an error
			CustomerCiry:    "",
		},
	}))
func (aa *accountAuthClient) ChangeAccount(ctx context.Context, in *basic.AccountInfoChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return aa.ChangeAccount(ctx, in, opts...)
}

*/
