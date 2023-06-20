package gatehttp

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandlerErrConnAA(err error, cancelCtxError context.CancelCauseFunc) string {
	st, _ := status.FromError(err)

	switch code := st.Code(); code {
	case codes.DeadlineExceeded:
		return "Deadline Exceeded"
	case codes.Canceled:
		return "Canceled"
	case codes.Internal:
		cancelCtxError(errors.New("INTERNAL:service-acct-aut"))
		return ""
	default:
		switch message := st.Message(); message {
		default:
			return message
		}
	}
}
