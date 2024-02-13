package servicedatacustomer

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

/*
var (
	ErrTokenFake          = jwt.ErrTokenFake
	ErrTokenExpired       = jwt.ErrTokenExpired
	ErrEmpty              = errors.New("empty value passed")
	ErrMissingMetadata    = errors.New("missing metadata")
	ErrMissingToken       = errors.New("missing token")
	ErrOrderStatNotSelect = errors.New("order status not selected")
)

func InterceptotHandlerErrors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	switch err {
	case nil:
		return resp, err
	case ErrTokenFake:
		return resp, status.New(codes.InvalidArgument, ErrTokenFake.Error()).Err()
	case ErrTokenExpired:
		return resp, status.New(codes.InvalidArgument, ErrTokenExpired.Error()).Err()
	case ErrEmpty:
		return resp, status.New(codes.InvalidArgument, ErrEmpty.Error()).Err()
	case ErrMissingMetadata:
		return resp, status.New(codes.InvalidArgument, ErrMissingMetadata.Error()).Err()
	case ErrMissingToken:
		return resp, status.New(codes.Unauthenticated, ErrMissingToken.Error()).Err()
	case ErrOrderStatNotSelect:
		return resp, status.New(codes.InvalidArgument, ErrOrderStatNotSelect.Error()).Err()
	default:
		return resp, status.New(codes.Internal, "").Err()
	}
}
*/
