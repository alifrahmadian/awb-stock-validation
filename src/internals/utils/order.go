package utils

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
