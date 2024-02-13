package apierror

import "errors"

var (
	ErrTokenFake          = errors.New("token fake")
	ErrTokenExpired       = errors.New("token expired")
	ErrEmpty              = errors.New("empty value passed")
	ErrMissingMetadata    = errors.New("missing metadata")
	ErrMissingToken       = errors.New("missing token")
	ErrOrderStatNotSelect = errors.New("order status not selected")
	ErrInvalidTokenFormat = errors.New("the heading Authorization must have the value: Bearer your_token")

	ErrIncorrectPass    = errors.New("incorrect password")
	ErrIncorrectLogin   = errors.New("incorrect login")
	ErrIncorrectCountry = errors.New("incorrect country")
	ErrPassLen          = errors.New("password is not in the allowed number of characters (8-64)")
	ErrLoginBusy        = errors.New("login busy")
	ErrDataLoss         = errors.New("data loss or corruption")
	ErrNotDataUpdate    = errors.New("data for update not provided")
)

// db error
var (
	ErrInvalidRequest   = errors.New("invalid request")
	ErrDeliveryLocation = errors.New("delivery country must match with warehouse side")
	ErrNotEnoughGoods   = errors.New("not enough goods")
	ErrNotEnoughMoney   = errors.New("not enough money")
	ErrNotCanceled      = errors.New("order completed cannot be canceled")
)

/*
	if _, ok := MapErr[err]; ok {
		anyone logics
	}
*/
var MapErr = map[error]bool{
	ErrTokenFake:          true,
	ErrTokenExpired:       true,
	ErrEmpty:              true,
	ErrMissingMetadata:    true,
	ErrMissingToken:       true,
	ErrOrderStatNotSelect: true,
	ErrInvalidTokenFormat: true,
	ErrIncorrectPass:      true,
	ErrIncorrectLogin:     true,
	ErrIncorrectCountry:   true,
	ErrPassLen:            true,
	ErrLoginBusy:          true,
	ErrDataLoss:           true,
	ErrInvalidRequest:     true,
	ErrDeliveryLocation:   true,
	ErrNotEnoughGoods:     true,
	ErrNotEnoughMoney:     true,
	ErrNotCanceled:        true,
}
