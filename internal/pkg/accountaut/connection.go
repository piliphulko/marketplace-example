package accountaut

import (
	pb "github.com/piliphulko/marketplace-example/internal/proto-genr/server-account-aut"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TakeConn(address string) (error, func(), ConnAccountAut) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err, nil, nil
	}

	return nil, func() { conn.Close() }, pb.NewAccountAutClient(conn)
}
