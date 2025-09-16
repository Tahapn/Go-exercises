package main

import (
	"fmt"
)

var balance int = 100

func main() {
	fmt.Println("Welcome. what do you want to do?")

	for {
		fmt.Println("1. Check balance")
		fmt.Println("2. add money")
		fmt.Println("3. withdraw money")
		fmt.Println("4. exit ")

		var choice int
		fmt.Print("enter the operation number:")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Print("Your balance is: ", balance)
		} else if choice == 2 {
			var number int
			fmt.Print("enter amount: ")
			fmt.Scan(&number)

			if number <= 0 {
				fmt.Print("The amount must be greater than 0.")
				continue
			}
			balance += number

		} else if choice == 3 {
			var number int
			fmt.Print("enter amount: ")
			fmt.Scan(&number)

			if number <= 0 || number > balance {
				fmt.Print("Invalid number. (<= 0 or greater than balance.)")
				continue
			}
			balance -= number
		} else {
			fmt.Println("Goodbye. ")
			break
		}
	}
	fmt.Println("Thanks!")
}

// Alternative way with switch

func alternativeWayWithSwithc() {
	fmt.Println("Welcome. What do you want to do?")

	for {
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. withdraw")
		fmt.Println("4. exit")

		var operation int
		fmt.Print("Enter the operation number: ")
		fmt.Scan(&operation)

		switch operation {
		case 1:
			fmt.Printf("Your current balance is: %d\n", balance)
		case 2:
			var depositAmount int
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Please try again.")
				continue
			}

			balance += depositAmount
		case 3:
			var withdrawalAmount int
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid withdrawal amount. Please try again.")
				continue
			}

			if withdrawalAmount > balance {
				fmt.Println("Insufficient funds. Please try again.")
				continue
			}

			balance -= withdrawalAmount
		case 4:
			fmt.Println("Goodbye.")
			return

		default:
			fmt.Println("Invalid operation. Please try again.")
		}
	}
}
