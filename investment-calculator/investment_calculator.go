package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 3.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)

	fmt.Printf("Future value of investment is %.2f\n", futureValue)
	fmt.Printf("Future real value with inflation of %.2f%% p.a. is %.2f\n", inflationRate, futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	return futureValue, futureRealValue
}
