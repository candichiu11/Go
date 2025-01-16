package main

import (
	"fmt"
)

func mainProfitCalculator() {
	// var revenue float64
	// var expenses float64
	// taxRate := 0.07
     
	revenue := getUserInput("Enter revenue: ")
	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)
	
    expenses := getUserInput("Enter expenses: ")
	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)
    
	taxRate := getUserInput("Enter tax rate: ")
	// fmt.Print("Enter tax rate: ")
	// fmt.Scan(&taxRate)
    
	profit, earningsBeforeTax, ratio := calculateProfit(revenue, expenses, taxRate)
    // earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate)
	// ratio := profit / revenue

	fmt.Printf("profit is $%.2f\n", profit)
	fmt.Printf("earningsBeforeTax is $%.2f\n", earningsBeforeTax)
	fmt.Printf("ratio is $%.2f\n", ratio)
}

func getUserInput(text string) float64{
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calculateProfit(revenue, expenses, taxRate float64) (profit float64, earningsBeforeTax float64, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - taxRate)
	ratio = profit / revenue
	return
}