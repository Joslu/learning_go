// Working with control structures

package main

import "fmt"

func main() {

	var accountBalanace float64 = 1000
	var moneyToDeposit float64
	var moneyToWithdraw float64

	userChoice := bank()

	if userChoice == 1 {
		fmt.Println("$", accountBalanace)
	} else if userChoice == 2 {
		fmt.Println("How much do you want to deposit?: ")
		fmt.Scan(&moneyToDeposit)

		if moneyToDeposit <= 0 {
			fmt.Println("Deposit should be greater tha $0")
			return
		}

		accountBalanace += moneyToDeposit // accountBalance = accountBalance + moneyToDeposit
		fmt.Println("This is your current balance: ", accountBalanace)
	} else if userChoice == 3 {
		fmt.Println("How much do you want to withdraw?: ")
		fmt.Scan(&moneyToWithdraw)

		if moneyToWithdraw <= 0 {
			fmt.Println("Value should be greater tha $0")
			return
		}

		if moneyToWithdraw > accountBalanace {
			print("You don't have enough money")
			return
		}

		accountBalanace -= moneyToWithdraw // accountBalance = accountBalance + moneyToDeposit
		fmt.Println("This is your current balance: ", accountBalanace)
	} else {
		fmt.Println("Bye 👋🏽 ")
	}
}

func bank() int {

	fmt.Println("Welcome to this GoBank 🏦 ")
	fmt.Println("What do you want to do today?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	return choice

}
