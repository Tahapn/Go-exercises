package main

import "fmt"

func main() {
	var revenue, expenses, taxes float64
	fmt.Print("enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("enter taxes: ")
	fmt.Scan(&taxes)

	var earningBeforeTax = revenue - expenses
	var earningAfterTax = (revenue - expenses) * (1 - taxes/100)
	var ratio = earningBeforeTax / earningAfterTax

	fmt.Println(earningBeforeTax)
	fmt.Println(earningAfterTax)
	fmt.Println(ratio)
}
