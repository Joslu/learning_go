// Working with control structures

package main

import "fmt"

func main() {

	var userChoice int
	var accountBalanace float64 = 1000
	var moneyToDeposit float64
	var moneyToWithdraw float64

	fmt.Println("Welcome to this GoBank 🏦 ")
	fmt.Println("What do you want to do today?")

	for {

		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			fmt.Println("$", accountBalanace)
		case 2:
			fmt.Println("How much do you want to deposit?: ")
			fmt.Scan(&moneyToDeposit)

			if moneyToDeposit <= 0 {
				fmt.Println("Deposit should be greater tha $0")
				//return
				continue
			}

			accountBalanace += moneyToDeposit // accountBalance = accountBalance + moneyToDeposit
			fmt.Println("This is your current balance: ", accountBalanace)
		case 3:
			fmt.Println("How much do you want to withdraw?: ")
			fmt.Scan(&moneyToWithdraw)

			if moneyToWithdraw <= 0 {
				fmt.Println("Value should be greater tha $0")
				//return
				continue
			}

			if moneyToWithdraw > accountBalanace {
				print("You don't have enough money")
				//return
				continue
			}

			accountBalanace -= moneyToWithdraw // accountBalance = accountBalance + moneyToDeposit
			fmt.Println("This is your current balance: ", accountBalanace)
		default:
			fmt.Println("Bye 👋🏽 ")
			fmt.Println("Thanks for choose our bank  ")
			return
		}
		/*
			if userChoice == 1 {
				fmt.Println("$", accountBalanace)
			} else if userChoice == 2 {
				fmt.Println("How much do you want to deposit?: ")
				fmt.Scan(&moneyToDeposit)

				if moneyToDeposit <= 0 {
					fmt.Println("Deposit should be greater tha $0")
					//return
					continue
				}

				accountBalanace += moneyToDeposit // accountBalance = accountBalance + moneyToDeposit
				fmt.Println("This is your current balance: ", accountBalanace)
			} else if userChoice == 3 {
				fmt.Println("How much do you want to withdraw?: ")
				fmt.Scan(&moneyToWithdraw)

				if moneyToWithdraw <= 0 {
					fmt.Println("Value should be greater tha $0")
					//return
					continue
				}

				if moneyToWithdraw > accountBalanace {
					print("You don't have enough money")
					//return
					continue
				}

				accountBalanace -= moneyToWithdraw // accountBalance = accountBalance + moneyToDeposit
				fmt.Println("This is your current balance: ", accountBalanace)
			} else {
				fmt.Println("Bye 👋🏽 ")
				break
			}
		*/
	}

}

func bank() int {
	var choice int
	for i := 0; i < 2; i++ {
		fmt.Println("Welcome to this GoBank 🏦 ")
		fmt.Println("What do you want to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

	}
	return choice
}
