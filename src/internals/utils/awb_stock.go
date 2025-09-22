package utils

import "strconv"

func ValidateAWBNumber(awbNumber string) bool {
	if len(awbNumber) != 12 {
		return false
	}

	if awbNumber[3] != '-' {
		return false
	}

	serialNumber := awbNumber[4:11]
	checkNumber := string(awbNumber[len(awbNumber)-1])

	serialNumberInt, err := strconv.Atoi(serialNumber)
	if err != nil {
		return false
	}

	checkNumberInt, err := strconv.Atoi(checkNumber)
	if err != nil {
		return false
	}

	return serialNumberInt%7 == checkNumberInt
}
