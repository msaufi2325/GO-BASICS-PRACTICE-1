package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//   => Show error message & exit if invalid input is provided
//   - No negative numbers
//   - Not 0
// 2) Store calculated results into file.

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func calculateEarnings(revenue, expenses, taxRate float64) (earningBeforeTax float64, earningAfterTax float64, ratio float64) {
	earningBeforeTax = revenue - expenses
	earningAfterTax = earningBeforeTax * (1 - taxRate/100)
	ratio = earningBeforeTax / earningAfterTax

	return earningBeforeTax, earningAfterTax, ratio
}

func storeResults(earningBeforeTax, earningAfterTax, ratio float64) {
	results := fmt.Sprintf("Earning Before Tax: %0.1f\nProfit: %0.1f\nRatio: %0.3f", earningBeforeTax, earningAfterTax, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------EXITING-----------------")
		return
		// panic(err)
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------EXITING-----------------")
		return
	}
	taxRate, err := getUserInput("Tax rate in %: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------EXITING-----------------")
		return
	}

	earningBeforeTax, earningAfterTax, ratio := calculateEarnings(revenue, expenses, taxRate)

	storeResults(earningBeforeTax, earningAfterTax, ratio)

	formattedEarningBeforeTax := fmt.Sprintf("Earning before tax: %0.1f\n", earningBeforeTax)
	formattedProfit := fmt.Sprintf("Profit: %0.1f\n", earningAfterTax)
	formattedRatio := fmt.Sprintf("Ratio: %0.1f\n", ratio)

	fmt.Print(formattedEarningBeforeTax, formattedProfit, formattedRatio)
}
