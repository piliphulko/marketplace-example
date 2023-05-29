package accountaut

import (
	"context"
	"fmt"

	"github.com/piliphulko/marketplace-example/internal/proto-genr/basic"
	pb "github.com/piliphulko/marketplace-example/internal/proto-genr/server-account-aut"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) AutAccount(ctx context.Context, loginPass *pb.LoginPass) (*pb.Reply, error) {
	if customerAut := loginPass.GetCustomerLoginPass(); customerAut != nil {
		fmt.Println(customerAut)
		return &pb.Reply{
			Reply: basic.REPLY_AUTHORIZED,
		}, status.New(codes.OK, "").Err()
	} else if warehouseAut := loginPass.GetWarehouseLoginPass(); warehouseAut != nil {
		fmt.Println(warehouseAut)
		return nil, nil
	} else if vendorAut := loginPass.GetVendorLoginPass(); vendorAut != nil {
		fmt.Println(vendorAut)
		return nil, nil
	} else {
		fmt.Println("error")
		return nil, nil
	}
}

func (s *server) CreateAccount(ctx context.Context, accountInfo *pb.AccountInfo) (*pb.Reply, error) {
	return nil, nil
}

func (s *server) UpdateAccount(ctx context.Context, accountInfo *pb.AccountInfo) (*pb.Reply, error) {
	return nil, nil
}
