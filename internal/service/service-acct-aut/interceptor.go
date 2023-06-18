package serviceacctaut

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InterceptotCheckCtx(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		return nil, status.New(codes.DeadlineExceeded, "").Err()
	case context.Canceled:
		return nil, status.New(codes.Canceled, "").Err()
	default:
		return handler(ctx, req)
	}
}
