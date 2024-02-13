package gatehttp

import (
	"bytes"
	"context"
	"fmt"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp/opt"
	"github.com/piliphulko/marketplace-example/internal/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func cutMessageFromGrpcAnswer(err error, cancelCtxError context.CancelCauseFunc) string {
	st, _ := status.FromError(err)

	switch code := st.Code(); code {
	case codes.DeadlineExceeded:
		return "Deadline Exceeded"
	case codes.Canceled:
		return "Canceled"
	case codes.Internal:
		cancelCtxError(err)
		return ""
	default:
		switch message := st.Message(); message {
		default:
			return message
		}
	}
}

func handlerErrGrpcApi(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, buf bytes.Buffer, err error) {
	if err != nil {
		if _, ok := apierror.MapErr[err]; ok {
			if errN := opt.WriteRedirectAnswerInfoErr(&buf, cutMessageFromGrpcAnswer(err, cancelCtxError)); errN != nil {
				cancelCtxError(fmt.Errorf("%s | %w", err.Error(), errN))
				return
			}
		}
		cancelCtxError(err)
		return
	}
}

func takeNickname(token string) (string, error) {
	jwtV, err := jwt.BeIntoJWT(token)
	if err != nil {
		return "", err
	}
	return jwtV.TakeNickname()
}
