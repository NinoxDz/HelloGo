package Libs

import (
	"errors"
	"fmt"
	"math"
)

func Sum(x int, y int) int {
	fmt.Println("___________sum function_____________")

	return x + y
}

func Sqrt(x float64) (float64, error) {
	fmt.Println("________multi return function_______")

	if x < 0 {
		return 0, errors.New("Undefined for negative value")
	}
	return math.Sqrt(x), nil
}
