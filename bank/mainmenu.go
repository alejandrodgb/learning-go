package main

import (
	"fmt"
)

func mainMenu() (userChoice int) {
	userChoice = 0

	for userChoice > 4 || userChoice < 1 {
		fmt.Println("Main Menu\nPlease select an option")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scan(&userChoice)

		fmt.Println()

		// Validate user input
		if userChoice > 4 || userChoice < 1 {
			fmt.Println("Invalid choice. Please try again")
		}
	}

	return userChoice
}
