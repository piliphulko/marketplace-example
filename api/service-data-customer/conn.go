package servicedatacustomer

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/service/service-data-customer/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type dataCustomerClient struct {
	core.DataCustomerClient
}

type DataCustomerClient interface {
	core.DataCustomerClient
}

type closeConn func()

func ConnToServiceDataCustomer(address string) (DataCustomerClient, closeConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return core.NewDataCustomerClient(conn), func() { conn.Close() }, nil
}

func (c *dataCustomerClient) GetCustomerInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*basic.CustomerInfo, error) {
	return c.GetCustomerInfo(ctx, in, opts...)
}

func (c *dataCustomerClient) GetCustomerOrders(ctx context.Context, in *basic.OrderStatus, opts ...grpc.CallOption) (*basic.Orders, error) {
	return c.GetCustomerOrders(ctx, in, opts...)
}

func (c *dataCustomerClient) GetWalletInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*basic.WalletInfo, error) {
	return c.GetWalletInfo(ctx, in, opts...)
}
