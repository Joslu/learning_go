/*
Build a profit calculator

- Ask for evenue, expenses, & tax rate
- Calculate earnings before taxe (EBT) and eranings after taxes (profit)
- Calculate ratio (EBT / profit)
- Ouput EBT, profit and the ratio
*/

package main

import (
	"fmt"
)

// Main function
func main() {

	var revenue, expenses, taxRate uint64

	fmt.Print("Type revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Type expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// EBT

	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

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
