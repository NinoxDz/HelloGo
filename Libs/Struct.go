package Libs

import "fmt"

func Struct() {
	fmt.Println("______________Struct________________")
	p := person{name: "Nadjmo", age: 28}
	fmt.Println(p)
	fmt.Println(p.age)
}

type person struct {
	name string
	age  int
}
