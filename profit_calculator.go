package main

import (
	"fmt"
)

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax rate in %: ")

	earningBeforeTax, earningAfterTax, ratio := calculateEarnings(revenue, expenses, taxRate)

	formattedEarningBeforeTax := fmt.Sprintf("Earning before tax: %0.1f\n", earningBeforeTax)
	formattedProfit := fmt.Sprintf("Profit: %0.1f\n", earningAfterTax)
	formattedRatio := fmt.Sprintf("Ratio: %0.1f\n", ratio)

	fmt.Print(formattedEarningBeforeTax, formattedProfit, formattedRatio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateEarnings(revenue, expenses, taxRate float64) (earningBeforeTax float64, earningAfterTax float64, ratio float64) {
	earningBeforeTax = revenue - expenses
	earningAfterTax = earningBeforeTax * (1 - taxRate/100)
	ratio = earningBeforeTax / earningAfterTax

	return earningBeforeTax, earningAfterTax, ratio
}
