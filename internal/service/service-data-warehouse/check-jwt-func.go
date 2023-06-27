package servicedatawarehouse

import (
	"context"
	"errors"

	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"google.golang.org/grpc/metadata"
)

func TakeLoginAndCheckJWTfromMetadataCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrMissingMetadata
	}
	JWT, err := jwt.BeIntoJWT(md.Get("authorization")[0])
	if err != nil {
		if errors.Is(err, jwt.ErrTokenFake) {
			return "", jwt.ErrTokenFake
		} else {
			return "", err
		}
	}
	if err := JWT.CheckJWT(); err != nil {
		if errors.Is(err, jwt.ErrTokenFake) {
			return "", jwt.ErrTokenFake
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return "", jwt.ErrTokenExpired
		} else {
			return "", err
		}
	}
	login, err := JWT.TakeNickname()
	if err != nil {
		return "", err
	}
	return login, nil
}
