package main

import "fmt"

func main() {
	revenue := getUserInput("Please Enter Revenue: ")
	expenses := getUserInput("Please Enter Expenses: ")
	taxRate := getUserInput("Please Enter Tax Rate: ")
	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Println("EBT :", ebt)
	fmt.Printf("future Value : %.1f\n", profit)
	fmt.Printf("RATIO : %.1f\n", ratio)
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoStr string) float64 {
	var userInput float64
	fmt.Print(infoStr)
	fmt.Scan(&userInput)
	return userInput
}
