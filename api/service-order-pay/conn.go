package serviceorderpay

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/service/service-order-pay/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type orderPayClient struct {
	core.OrderPayClient
}

type OrderPayClient interface {
	core.OrderPayClient
}

type closeConn func()

func ConnToServiceOrderPay(address string) (OrderPayClient, closeConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return core.NewOrderPayClient(conn), func() { conn.Close() }, nil
}

func (c *orderPayClient) CreateOrder(ctx context.Context, in *basic.NewOrderARRAY, opts ...grpc.CallOption) (*basic.OrderUuid, error) {
	return c.CreateOrder(ctx, in, opts...)
}

func (c *orderPayClient) ConfirmOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.ConfirmOrder(ctx, in, opts...)
}

func (c *orderPayClient) CancelOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.CancelOrder(ctx, in, opts...)
}

func (c *orderPayClient) CompleteOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.CompleteOrder(ctx, in, opts...)
}

func (c *orderPayClient) AddToOrder(ctx context.Context, in *basic.AddToOrderARRAY, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.AddToOrder(ctx, in, opts...)
}
