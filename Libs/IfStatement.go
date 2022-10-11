package Libs

import "fmt"

func If(age int) {
	fmt.Println("___________if statement_____________")

	//if statement
	if age > 25 {
		fmt.Println("you are older then 25")
	} else if age < 50 {
		fmt.Println("you are younger then 50")
	} else {
		fmt.Println("you are older then 50")

	}
}
