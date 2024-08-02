/*
Build a profit calculator

- Ask for evenue, expenses, & tax rate
- Calculate earnings before taxe (EBT) and eranings after taxes (profit)
- Calculate ratio (EBT / profit)
- Ouput EBT, profit and the ratio
*/

/*

	UPDATE:

	Goals

	1) Validate user input
		- Show error messages and exit if invalid input is provided
		- no negative numbers
		- not 0

	2) Stored calculated values into file

*/

package main

import (
	"errors"
	"fmt"
	"os"
)

const profitValue = "profit.txt"

func getUserInput(infoText string) (float64, error) {
	var userinput float64

	fmt.Println(infoText)
	fmt.Scan(&userinput)

	if userinput <= 0 {

		return 0, errors.New("Value must be greater that 0")
	}

	return userinput, nil

}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit

	return ebt, profit, ratio

}

func saveValues(valueToSave1, valueToSave2, valueToSave3 float64) {

	valueText := fmt.Sprintf("EBT:%.1f\nProfit:%.1f\nsRatio:%.1f\n", valueToSave1, valueToSave2, valueToSave3)
	os.WriteFile("value.txt", []byte(valueText), 0644)

}

// Main function
func main() {

	fmt.Println("Welcome to ProfitCalculator 🧮")

	revenue, err1 := getUserInput("Revenue: ")

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	expenses, err2 := getUserInput("Expenses: ")

	if err2 != nil {
		fmt.Println(err2)
		return
	}
	taxRate, err3 := getUserInput("Tax Rate")

	if err3 != nil {
		fmt.Println(err3)
		return
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Println("ebt: ", ebt)
	fmt.Println("profit: ", profit)
	fmt.Println("ratio: ", ratio)

	saveValues(ebt, profit, ratio)
}

func getEbt() {

	fmt.Print("GET ebt FUNCTION")

}

func getProfit() {
	fmt.Print("GET Profit FUNCTION")
}

func getRatio() {
	fmt.Print("GET Ratio FUNCTION")
}
