package errors

import "errors"

var (
	ErrOrderAWBRequired         = errors.New("awb required")
	ErrOrderSenderRequired      = errors.New("sender required")
	ErrOrderReceiverRequired    = errors.New("receiver required")
	ErrOrderTotalWeightRequired = errors.New("total weight required")
)
