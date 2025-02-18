package main

import (
	"fmt"

	"github.com/alejandrodgb/learning-go/fileops"
)

func main() {

	file := "balance.txt"

	accountBalance, err := fileops.ReadFloatFromFile(file)

	if err != nil {
		fmt.Println(err, "Starting with $0 balance")
		fileops.WriteFloatToFile(0, file)
	}

	fmt.Println("Welcome to GoBank")

	keepGoing := true

	for keepGoing {
		userChoice := mainMenu()

		switch userChoice {
		case 1:
			fmt.Printf("Your account balance is: $%.2f\n", accountBalance)
			fmt.Printf("Returning to main menu\n\n")
		case 2:
			deposit := depositMoney()
			accountBalance += deposit
			fileops.WriteFloatToFile(accountBalance, file)
			fmt.Printf("You account balance after depositing $%.2f is: $%.2f\n", deposit, accountBalance)
			fmt.Printf("Returning to main menu\n\n")
		case 3:
			withdrawal := withdrawMoney(accountBalance)
			accountBalance -= withdrawal
			fileops.WriteFloatToFile(accountBalance, file)
			fmt.Printf("You account balance after withdrawing $%.2f is: $%.2f\n", withdrawal, accountBalance)
			fmt.Printf("Returning to main menu\n\n")
		default:
			keepGoing = false
			fmt.Println("Thank you for using GoBank")
		}
	}
}

func depositMoney() (deposit float64) {

	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&deposit)

	if deposit < 0 {
		fmt.Println("Invalid amount")
		return 0
	}

	return deposit
}

func withdrawMoney(balance float64) (withdrawal float64) {

	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&withdrawal)

	if withdrawal < 0 || withdrawal > balance {
		fmt.Println("Invalid amount. $0 have been widthrawn")
		return 0
	}

	return withdrawal
}
