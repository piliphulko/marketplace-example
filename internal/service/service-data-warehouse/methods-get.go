package servicedatawarehouse

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/basic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) GetAcctInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*basic.WarehouseInfo, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	token := md.Get("authorization")[0]
}
