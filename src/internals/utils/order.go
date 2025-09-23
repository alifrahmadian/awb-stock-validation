package utils

import (
	"github.com/audricimanuel/awb-stock-allocation/utils/constants"
	e "github.com/audricimanuel/awb-stock-allocation/utils/errors"
)

func CalculateTotalPrice(totalWeight float64) int {
	if totalWeight >= 0 && totalWeight < 10 {
		return int(totalWeight) * 5000
	} else if totalWeight >= 10 && totalWeight < 20 {
		return int(totalWeight) * 4500
	} else if totalWeight >= 20 && totalWeight < 25 {
		return int(totalWeight) * 4000
	} else {
		return int(totalWeight) * 3500
	}
}

func MapInputStatusToString(input int) (string, error) {
	switch input {
	case constants.ORDERINPUT_STATUS_PENDING:
		return constants.ORDER_STATUS_PENDING, nil
	case constants.ORDERINPUT_STATUS_CONFIRM:
		return constants.ORDER_STATUS_CONFIRM, nil
	case constants.ORDERINPUT_STATUS_SHIPPED:
		return constants.ORDER_STATUS_SHIPPED, nil
	case constants.ORDERINPUT_STATUS_COMPLETED:
		return constants.ORDER_STATUS_COMPLETED, nil
	case constants.ORDERINPUT_STATUS_CANCELLED:
		return constants.ORDER_STATUS_CANCELLED, nil
	default:
		return "", e.ErrOrderStatusInvalid
	}
}
