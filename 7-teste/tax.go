package tax

func CalculateTaxt(amount float64) float64 {
	if amount >= 100 {
		return 10.0
	}
	return 5
}