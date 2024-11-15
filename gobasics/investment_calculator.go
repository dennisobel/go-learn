package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	// years, expectedReturnRate :=  10.0, 5.5

	outputText("Investment Amount: ")

	// fmt.Print("Enter the initial investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Number of years: ")
	fmt.Scan(&years)

	outputText("Investment Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount,expectedReturnRate,inflationRate,years)

	formatedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formatedRFV := fmt.Sprintf("Real Future Value: %.2f\n", futureRealValue)
	fmt.Println("Fututre Value: ", futureValue)
	fmt.Print(formatedFV)
	fmt.Print(formatedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}
