package Libs

import "fmt"

func Slice() {
	fmt.Println("____________slice____________")
	/// slice

	sliceArray := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sliceArray)
	sliceArray = append(sliceArray, 7)
	fmt.Println(sliceArray)
}
