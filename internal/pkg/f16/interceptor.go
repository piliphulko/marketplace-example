package f16

import (
	"context"
	"fmt"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var mapClientCodes = map[error]codes.Code{
	apierror.ErrTokenFake:          codes.InvalidArgument,
	apierror.ErrTokenExpired:       codes.InvalidArgument,
	apierror.ErrEmpty:              codes.InvalidArgument,
	apierror.ErrMissingMetadata:    codes.InvalidArgument,
	apierror.ErrMissingToken:       codes.Unauthenticated,
	apierror.ErrOrderStatNotSelect: codes.InvalidArgument,
	apierror.ErrPassLen:            codes.InvalidArgument,
	apierror.ErrIncorrectPass:      codes.InvalidArgument,
	apierror.ErrIncorrectLogin:     codes.Unauthenticated,
	apierror.ErrDataLoss:           codes.DataLoss,
	apierror.ErrLoginBusy:          codes.AlreadyExists,
	apierror.ErrIncorrectCountry:   codes.InvalidArgument,
	apierror.ErrInvalidRequest:     codes.InvalidArgument,
	apierror.ErrNotEnoughGoods:     codes.InvalidArgument,
	apierror.ErrDeliveryLocation:   codes.InvalidArgument,
	apierror.ErrNotCanceled:        codes.InvalidArgument,
	apierror.ErrNotEnoughMoney:     codes.InvalidArgument,
	apierror.ErrInvalidTokenFormat: codes.InvalidArgument,
}

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

func IntrceptorHandlerErrors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err == nil {
		return resp, err
	}
	fmt.Println(err)
	codesAnswer, ok := mapClientCodes[err]
	if ok {
		return resp, status.New(codesAnswer, err.Error()).Err()
	} else {
		return resp, status.New(codes.Internal, "").Err()
	}
	/*
		switch err {
		case nil:
			return resp, err
		case apierror.ErrTokenFake:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrTokenFake.Error()).Err()
		case apierror.ErrTokenExpired:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrTokenExpired.Error()).Err()
		case apierror.ErrEmpty:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrEmpty.Error()).Err()
		case apierror.ErrMissingMetadata:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrMissingMetadata.Error()).Err()
		case apierror.ErrMissingToken:
			fmt.Println(err)
			return resp, status.New(codes.Unauthenticated, apierror.ErrMissingToken.Error()).Err()
		case apierror.ErrOrderStatNotSelect:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrOrderStatNotSelect.Error()).Err()
		case apierror.ErrPassLen:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrPassLen.Error()).Err()
		case apierror.ErrIncorrectPass:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrIncorrectPass.Error()).Err()
		case apierror.ErrIncorrectLogin:
			fmt.Println(err)
			return resp, status.New(codes.Unauthenticated, apierror.ErrIncorrectLogin.Error()).Err()
		case apierror.ErrDataLoss:
			fmt.Println(err)
			return resp, status.New(codes.DataLoss, apierror.ErrDataLoss.Error()).Err()
		case apierror.ErrLoginBusy:
			fmt.Println(err)
			return resp, status.New(codes.AlreadyExists, apierror.ErrLoginBusy.Error()).Err()
		case apierror.ErrIncorrectCountry:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrIncorrectCountry.Error()).Err()
		case apierror.ErrInvalidRequest:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrInvalidRequest.Error()).Err()
		case apierror.ErrNotEnoughGoods:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrNotEnoughGoods.Error()).Err()
		case apierror.ErrDeliveryLocation:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrDeliveryLocation.Error()).Err()
		case apierror.ErrNotCanceled:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrNotCanceled.Error()).Err()
		case apierror.ErrNotEnoughMoney:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrNotEnoughMoney.Error()).Err()
		case apierror.ErrInvalidTokenFormat:
			fmt.Println(err)
			return resp, status.New(codes.InvalidArgument, apierror.ErrInvalidTokenFormat.Error()).Err()
		default:
			fmt.Println(err)
			return resp, status.New(codes.Internal, "").Err()
		}
	*/
}
