package main

import (
	"fmt"
	"math"
)

// Constant
const inflation float64 = 3.4

func main() {

	var investmentAmount float64
	var expectedReturnDate float64
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return date: ")
	fmt.Scan(&expectedReturnDate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	//futureValue := investmentAmount * math.Pow((1+expectedReturnDate/100), years)
	//futureRealValue := futureValue / math.Pow(1+inflation, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnDate, years)

	fmt.Println(futureValue)

	// Formating Strings

	//formatedFV := fmt.Sprintf("Future value: %.1f \n", futureValue)
	//formatedRFV := fmt.Sprintf("Future real value %.1f \n", futureRealValue)

	//fmt.Print(formatedFV, formatedRFV)

	fmt.Println("Future real value: ", futureRealValue)
	fmt.Printf("Future real value %v \n", futureRealValue)

	fmt.Printf(`Future value: %v |
	Future real value %v \n`, futureValue, futureRealValue)

	fmt.Printf(`Future value: %.1f | 
	Future real value %.1f \n`, futureValue, futureRealValue)
}

func outputText(textToWrite string) {

	fmt.Print(textToWrite)

}

// returning values
func calculateFutureValues(invAmount float64, expReturnDate float64, years float64) (float64, float64) {
	fv := invAmount * math.Pow((1+expReturnDate/100), years)
	frV := fv / math.Pow(1+inflation, years)

	return fv, frV
}

func calculateFutureValuesAlternative(invAmount float64, expReturnDate float64, years float64) (fv float64, frV float64) {
	fv = invAmount * math.Pow((1+expReturnDate/100), years)
	frV = fv / math.Pow(1+inflation, years)

	//return fv, frV
	return
}
