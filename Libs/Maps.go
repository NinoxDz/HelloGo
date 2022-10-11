package Libs

import "fmt"

func Map() {
	fmt.Println("________________MAP_________________")
	//// map

	varMap := make(map[string]int)
	varMap["triangle"] = 2
	varMap["square"] = 3

	fmt.Println(varMap)
	fmt.Println(varMap["square"])
	delete(varMap, "square")
	fmt.Println(varMap)
}
