package main

import (
	"errors"
	"fmt"
)

// Goals
// 1) Validate user input
//   => Show error message & exit if invalid input is provided
//   - No negative numbers
//   - Not 0
// 2) Store calculated results into file.

func main() {

	revenue, err := getUserInput("Revenue: ")
	expenses, err := getUserInput("Expenses: ")
	taxRate, err := getUserInput("Tax rate in %: ")

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("---------------------------------")
	}

	earningBeforeTax, earningAfterTax, ratio := calculateEarnings(revenue, expenses, taxRate)

	formattedEarningBeforeTax := fmt.Sprintf("Earning before tax: %0.1f\n", earningBeforeTax)
	formattedProfit := fmt.Sprintf("Profit: %0.1f\n", earningAfterTax)
	formattedRatio := fmt.Sprintf("Ratio: %0.1f\n", ratio)

	fmt.Print(formattedEarningBeforeTax, formattedProfit, formattedRatio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("invalid amount must be greater than 0")
	}

	return userInput, nil
}

func calculateEarnings(revenue, expenses, taxRate float64) (earningBeforeTax float64, earningAfterTax float64, ratio float64) {
	earningBeforeTax = revenue - expenses
	earningAfterTax = earningBeforeTax * (1 - taxRate/100)
	ratio = earningBeforeTax / earningAfterTax

	return earningBeforeTax, earningAfterTax, ratio
}
