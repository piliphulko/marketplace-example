package servicedatawarehouse

import (
	"context"

	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InterceptorCheckCtx(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		return nil, status.New(codes.DeadlineExceeded, "").Err()
	case context.Canceled:
		return nil, status.New(codes.Canceled, "").Err()
	default:
		return handler(ctx, req)
	}
}

func InterceptorHandlerError(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		switch err {
		case jwt.ErrTokenFake:
			return resp, status.Error(codes.Unauthenticated, jwt.ErrTokenFake.Error())
		case jwt.ErrTokenExpired:
			return resp, status.Error(codes.Unauthenticated, jwt.ErrTokenExpired.Error())
		case ErrMissingMetadata:
			return resp, status.Error(codes.DataLoss, ErrMissingMetadata.Error())
		default:
			return resp, status.Error(codes.Internal, "")
		}
	}
	return resp, status.New(codes.OK, "").Err()
}
