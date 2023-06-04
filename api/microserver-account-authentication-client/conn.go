package microserveraccountauthenticationclient

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/service/microserver-account-authentication/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type closeConn func()

type AccountAutClient interface {
	core.AccountAutClient
}

type accountAutClient struct {
	core.AccountAutClient
}

func ConnToMicroserverAccountAuthentication(address string) (AccountAutClient, closeConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return core.NewAccountAutClient(conn), func() { conn.Close() }, nil
}

func (aa *accountAutClient) AutAccount(ctx context.Context, in *basic.LoginPass, opts ...grpc.CallOption) (*basic.Reply, error) {
	return aa.AutAccount(ctx, in, opts...)
}

func (aa *accountAutClient) CreateAccount(ctx context.Context, in *basic.AccountInfo, opts ...grpc.CallOption) (*basic.Reply, error) {
	return aa.CreateAccount(ctx, in, opts...)
}

func (aa *accountAutClient) UpdateAccount(ctx context.Context, in *basic.AccountInfo, opts ...grpc.CallOption) (*basic.Reply, error) {
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
