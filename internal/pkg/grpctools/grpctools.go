package grpctools

import (
	"context"
	"errors"
	"strings"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"google.golang.org/grpc/metadata"
)

func TakeLoginAndCheckJWT(token string) (string, error) {
	if token == "" {
		return "", apierror.ErrEmpty
	}
	JWT, err := jwt.BeIntoJWT(token)
	if err != nil {
		if errors.Is(err, apierror.ErrTokenFake) {
			return "", apierror.ErrTokenFake
		} else {
			return "", err
		}
	}
	if err := JWT.CheckJWT(); err != nil {
		if errors.Is(err, apierror.ErrTokenFake) {
			return "", apierror.ErrTokenFake
		} else if errors.Is(err, apierror.ErrTokenExpired) {
			return "", apierror.ErrTokenExpired
		} else {
			return "", err
		}
	}
	return JWT.TakeNickname()
}

func TakeJWTfromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", apierror.ErrMissingMetadata
	}
	token, ok := md["authorization"]
	if !ok {
		return "", apierror.ErrMissingToken
	}
	return strings.TrimPrefix(token[0], "Bearer "), nil
}
