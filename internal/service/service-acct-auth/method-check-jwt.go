package serviceacctauth

import (
	"context"
	"errors"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/api/basic"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CheckJWT(ctx context.Context, in *basic.StringJWT) (*emptypb.Empty, error) {

	// CHECK EMPTY
	if in == nil && &in.StringJwt == nil && in.StringJwt == "" {
		return &emptypb.Empty{}, apierror.ErrEmpty
	}
	JWT, err := jwt.BeIntoJWT(in.StringJwt)
	if err != nil {
		if errors.Is(err, apierror.ErrTokenFake) {
			return &emptypb.Empty{}, apierror.ErrTokenFake
		} else {
			return &emptypb.Empty{}, err
		}
	}
	if err := JWT.CheckJWT(); err != nil {
		if errors.Is(err, apierror.ErrTokenFake) {
			return &emptypb.Empty{}, apierror.ErrTokenFake
		} else if errors.Is(err, apierror.ErrTokenExpired) {
			return &emptypb.Empty{}, apierror.ErrTokenExpired
		} else {
			return &emptypb.Empty{}, err
		}
	}
	return &emptypb.Empty{}, status.New(codes.OK, "").Err()
}
