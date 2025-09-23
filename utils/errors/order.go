package errors

import "errors"

var (
	ErrOrderAWBRequired             = errors.New("awb required")
	ErrOrderSenderRequired          = errors.New("sender required")
	ErrOrderReceiverRequired        = errors.New("receiver required")
	ErrOrderTotalWeightRequired     = errors.New("total weight required")
	ErrOrderNotFound                = errors.New("order not found")
	ErrOrderStatusPendingValidation = errors.New("pending order can only be confirmed or cancelled")
	ErrOrderStatusConfirmValidation = errors.New("confirm order can only be shipped or cancelled")
	ErrOrderStatusShippedValidation = errors.New("shipped order can only be completed")
	ErrOrderStatusFinal             = errors.New("unable to update the order status")
	ErrOrderStatusInvalid           = errors.New("invalid order status")
)
