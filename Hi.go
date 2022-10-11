package main

import (
	"HelloGo/Libs"
	"fmt"
)

func main() {
	//print
	fmt.Println("Hello this is Go")

	var name string = "nadjmo"
	age := 28
	fmt.Println(name, age)

	Libs.If(age)
	Libs.Array()
	Libs.Slice()
	Libs.ForLoop()
	Libs.WhileLoop()
	Libs.ForeachLoop()
	Libs.Map()
	Libs.Pointer()
	Libs.Struct()

	sumResult := Libs.Sum(5, 4)
	fmt.Println(sumResult)

	sqrtResult, sqrtError := Libs.Sqrt(16)
	fmt.Println("sqrt of", 16, sqrtResult)

	varSqrt := 16.5
	sqrtResult, sqrtError = Libs.Sqrt(varSqrt)

	if sqrtError != nil {
		fmt.Println("error sqrt of:", varSqrt, sqrtError)
	} else {
		fmt.Println("sqrt of", varSqrt, sqrtResult)
	}
	varSqrt = -16
	sqrtResult, sqrtError = Libs.Sqrt(varSqrt)

	if sqrtError != nil {
		fmt.Println("error sqrt of:", varSqrt, sqrtError)
	} else {
		fmt.Println("sqrt of", varSqrt, sqrtResult)
	}

}
