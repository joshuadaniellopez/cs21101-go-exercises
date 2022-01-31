package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	if successRate <= 100 && successRate >= 0 {
		return (float64(productionRate) * (successRate / 100))
	} else {
		return 0
	}
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	if successRate <= 100 && successRate >= 0 {
		return int((float64(productionRate) * (successRate / 100)) / 60)
	} else {
		return 0
	}
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	batchOfTen := carsCount / 10
	remainderOfTen := carsCount % 10

	return uint(batchOfTen)*95000 + uint(remainderOfTen)*10000
}
