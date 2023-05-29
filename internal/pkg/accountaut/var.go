package accountaut

import (
	pb "github.com/piliphulko/marketplace-example/internal/proto-genr/server-account-aut"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
)

var (
	LogGRPC *zapgrpc.Logger
)

type ConnAccountAut = pb.AccountAutClient

type server struct {
	pb.UnimplementedAccountAutServer
}

func NewServer() *server { return &server{} }

func RegisterSever(s1 *grpc.Server, s2 *server) {
	pb.RegisterAccountAutServer(s1, s2)
}
