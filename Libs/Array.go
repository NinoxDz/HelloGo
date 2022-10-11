package Libs

import "fmt"

func Array() {
	fmt.Println("_____Array Example_____")

	//fixed array
	var array [5]int
	array[2] = 2
	fmt.Println(array)

	arrayTwo := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayTwo)

}
