package main

import (
	"fmt"
	"math"
)

func main() {

	// Constant
	const inflation float64 = 3.4

	var investmentAmount float64
	var expectedReturnDate float64
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return date: ")
	fmt.Scan(&expectedReturnDate)

	fmt.Scan("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnDate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflation, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
