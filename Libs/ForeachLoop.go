package Libs

import "fmt"

func ForeachLoop() {
	fmt.Println("___________Foreach Loop_____________")

	loopArr := []string{"a", "b", "c", "d"}

	for index, value := range loopArr {
		fmt.Print(index, "=>", value, "\n")
	}
}
