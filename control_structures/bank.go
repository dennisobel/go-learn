package main

import (
	"fmt"
	"example.com/control-structures/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"



func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println(("ERROR"))
		fmt.Println(err)
		fmt.Println("--------")
		// panic("Can't continue without the balance file!")
	}
	fmt.Println("WELCOME TO GO BANK")
	fmt.Println("REACH US 247 @", randomdata.PhoneNumber())
	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Print("Your deposit.")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount, must be greater than zero")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your balance has been updated! New Amount:  ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdrawal.")
			var withdrawAount float64
			fmt.Scan((&withdrawAount))

			if withdrawAount <= 0 {
				fmt.Println("Invalid amount, must be greater than zero")
				continue
			}

			if withdrawAount > accountBalance {
				fmt.Println("Invalid amount, you can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawAount
			fmt.Println("Your balance has been updated! New Amount:  ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye")
			// break
			return
		}
	}
}
