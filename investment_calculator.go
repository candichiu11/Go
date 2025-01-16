package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

//	fmt.Print("Enter investment amount: ")
    outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

//	fmt.Print("Enter number of years: ")
	outputText("Enter number of years: ")
	fmt.Scan(&years)

//	fmt.Print("Enter expected return rate: ")
	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
    
	// \n is a newline character
	// fmt.Printf("Future value is $%.2f\n", futureValue) 
	// fmt.Printf("Future value (adjusted for inflation) is $%.2f\n", futureRealValue)
	//fmt.Printf("Future value is %v\nFuture value (adjusted for inflation) is %v\n", futureValue, futureRealValue)

	formattedFV := fmt.Sprintf("Future value is %.0f\n", futureValue)
	formattedFVInflation := fmt.Sprintf("Future value (adjusted for inflation) is %.0f\n", futureRealValue)
	
	fmt.Print(formattedFV, formattedFVInflation)

//	mainProfitCalculator()

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)
	//return fv, rfv
	return
}
