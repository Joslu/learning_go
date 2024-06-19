package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount float64 = 1000
	var expectedReturnDate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow((1+expectedReturnDate/100), years)

	fmt.Println(futureValue)
}
