package accountaut

import (
	"context"
	"fmt"

	pb "github.com/piliphulko/marketplace-example/internal/proto-genr/server-account-aut"
)

type server struct {
	pb.UnimplementedAccountAutServer
}

func (s *server) AutAccount(ctx context.Context, loginPass *pb.LoginPass) (*pb.Reply, error) {
	if customerAut := loginPass.GetCustomerLoginPass(); customerAut != nil {
		fmt.Println()
		return nil, nil
	} else if warehouseAut := loginPass.GetWarehouseLoginPass(); warehouseAut != nil {
		fmt.Println()
		return nil, nil
	} else if vendorAut := loginPass.GetVendorLoginPass(); vendorAut != nil {
		fmt.Println()
		return nil, nil
	} else {
		fmt.Println()
		return nil, nil
	}
}

func (s *server) CreateAccount(ctx context.Context, accountInfo *pb.AccountInfo) (*pb.Reply, error) {
	return nil, nil
}

func (s *server) UpdateAccount(ctx context.Context, accountInfo *pb.AccountInfo) (*pb.Reply, error) {
	return nil, nil
}
