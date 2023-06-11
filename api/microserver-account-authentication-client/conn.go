package microserveraccountauthenticationclient

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/service/microserver-account-authentication/core"
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

func ConnToMicroserverAccountAuthentication(address string) (AccountAutClient, closeConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return core.NewAccountAutClient(conn), func() { conn.Close() }, nil
}

func (aa *accountAutClient) AutAccount(ctx context.Context, in *basic.LoginPass, opts ...grpc.CallOption) (*basic.StringJWT, error) {
	return aa.AutAccount(ctx, in, opts...)
}

func (aa *accountAutClient) CheckJWT(ctx context.Context, in *basic.StringJWT, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return aa.CheckJWT(ctx, in, opts...)
}

func (aa *accountAutClient) CreateAccount(ctx context.Context, in *basic.AccountInfoChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return aa.CreateAccount(ctx, in, opts...)
}

func (aa *accountAutClient) UpdateAccount(ctx context.Context, in *basic.AccountInfoChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return aa.UpdateAccount(ctx, in, opts...)
}

/*
reply, err := a.AutAccount(context.Background(), &basic.LoginPass{
		AccountChoice: &basic.LoginPass_CustomerLoginPass{
			CustomerLoginPass: &basic.CustomerAut{
				LoginCustomer:    "123",
				PasswortCustomer: "211",
			},
		},
	})
*/
