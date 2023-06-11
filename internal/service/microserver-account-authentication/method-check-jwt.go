package microserveraccountauthentication

import (
	"context"
	"errors"

	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CheckJWT(ctx context.Context, in *basic.StringJWT) (*emptypb.Empty, error) {
	JWT, err := jwt.BeIntoJWT(in.StringJwt)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenFake) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, jwt.ErrTokenFake.Error()).Err()
		} else {
			return &emptypb.Empty{}, status.New(codes.Internal, "").Err()
		}
	}
	if err := JWT.CheckJWT(); err != nil {
		if errors.Is(err, jwt.ErrTokenFake) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, jwt.ErrTokenFake.Error()).Err()
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return &emptypb.Empty{}, status.New(codes.Unauthenticated, jwt.ErrTokenExpired.Error()).Err()
		}
	}
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}
