// Working with control structures

package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanaceFileName = "balance.txt"

func main() {

	var userChoice int

	var moneyToDeposit float64
	var moneyToWithdraw float64

	accountBalanace, err := fileops.GetFloatFromFile(accountBalanaceFileName)

	if err != nil {
		fmt.Println("ERROR 😰: ")
		fmt.Println(err)
		fmt.Println("---------------------")

		//return
		panic("Can't continue sorry")
	}

	fmt.Println("Welcome to this GoBank 🏦 ")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())
	fmt.Println("What do you want to do today?")

	for {

		presenOptions()

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
			fileops.WriteFloatToFile(accountBalanace, accountBalanaceFileName)
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
			fileops.WriteFloatToFile(accountBalanace, accountBalanaceFileName)
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
