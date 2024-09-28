package main

// Goals
// 1) Validate user input
// => Show error message & exit if invalid input is provided
// - No negative numbers
// - No 0
// 2) Store calculated results into file

import (
	"fmt"
	"os"
)

func userInputData(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 {
		panic("The value must be greater than 0.")
	}

	return value
}

func main() {
	var revenue, expenses, taxes float64

	revenue = userInputData("Enter revenue: ")
	expenses = userInputData("Enter expenses: ")
	taxes = userInputData("Enter taxes: ")

	var earningBeforeTax = revenue - expenses
	var earningAfterTax = (revenue - expenses) * (1 - taxes/100)
	var ratio = earningBeforeTax / earningAfterTax

	var output string = fmt.Sprintf("Earning before tax: %v.\nAfter tax %v.\nRatio:%.2f", earningBeforeTax, earningAfterTax, ratio)

	fmt.Println(output)

	os.WriteFile("output.txt", []byte(output), 0644)
}
