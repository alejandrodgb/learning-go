package main

import (
	"fmt"
	"os"
)

type Expense struct {
	Name   string
	Amount float64
}

func main() {

	var revenue float64 = userInput("Enter revenue: ")
	var taxRate float64 = userInput("Enter tax rate: ")

	expenses := []Expense{}

	for {
		name, amount := getExpenseDetails()

		expenses = append(expenses, Expense{name, amount})

		var addMore string
		fmt.Print("Add more expenses? (yes/no): ")
		fmt.Scan(&addMore)

		if addMore != "yes" {
			break
		}
	}

	ebt, profit, ratio, totalExpenses := financialSummary(revenue, taxRate, expenses)

	writeToFile(fmt.Sprintf("%.2f\n%.2f\n%.2f\n%.2f\n%.2f", revenue, totalExpenses, ebt, profit, ratio))

	fmt.Println("Revenue: ", revenue)
	fmt.Println("Total Expenses: ", totalExpenses)
	fmt.Println("Expenses: ", expenses)
	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)

}

func userInput(message string) (input float64) {
	fmt.Print(message)
	fmt.Scan(&input)

	if input <= 0 {
		fmt.Println("Invalid input. Please try again")
		panic("Invalid input")
	}

	return input
}

func financialSummary(revenue float64, taxRate float64, expenses []Expense) (ebt, profit, ratio, sumOfExpenses float64) {

	for _, expense := range expenses {
		sumOfExpenses += expense.Amount
	}

	ebt = revenue - sumOfExpenses
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt / profit

	return ebt, profit, ratio, sumOfExpenses
}

func getExpenseDetails() (string, float64) {
	var name string
	var amount float64

	fmt.Print("Enter expense name: ")
	fmt.Scan(&name)
	fmt.Print("Enter expense amount: ")
	fmt.Scan(&amount)

	return name, amount
}

func writeToFile(stringToWrite string) {
	os.WriteFile("financial_summary.txt", []byte(fmt.Sprint(stringToWrite)), 0644)
}
