package Libs

import "fmt"

func Pointer() {
	fmt.Println("______________pointer_______________")

	pointer := 7
	fmt.Println(pointer)
	fmt.Println(&pointer)
	inc(&pointer)
	fmt.Println(pointer)
}

func inc(x *int) {
	*x++
}
