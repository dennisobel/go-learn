package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResults(ebt float64) {
	results := fmt.Sprintf("EBT: %1.f\n", ebt)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func main() {
	revenue, err1 := getInput("ENTER REVENUE: ")

	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	expenses, err2 := getInput("ENTER EXPENSES: ")

	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return
	// }

	taxRate, err3 := getInput("ENTER TAX RATE: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt := calculateEBT(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profit)

	storeResults(ebt)

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Profit Ratio:", ratio)
}

func getInput(prompt string) (float64, error) {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("Invalid Input")
	}

	return input, nil
}

func calculateEBT(revenue, expenses float64) float64 {
	return revenue - expenses
}

func calculateProfit(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}
