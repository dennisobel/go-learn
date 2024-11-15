package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	// appUser, err := newUser(userFirstName, userLastName, userBirthDate)
	appUser,err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println((err))
		return
	}

	admin := user.NewAdmin("test@test.com","password")
	admin.OutputUserDetails()
	admin.ClearUserName()	

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

