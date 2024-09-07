package main

import (
	"fmt"

	"example.com/starting_project/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	admin := user.NewAdmin("a@email.com", "20304")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {

		fmt.Println(err)
		return

	}

	// ... do something awesome with that gathered data!
	appUser.Data()
	appUser.ClearName()
	appUser.Data()

	admin.User.Data()
	admin.User.ClearName()
	admin.User.Data()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
