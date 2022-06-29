package main

import "fmt"

func ArrayDataType() {
	var names []string
	var values = []int {
		1,2,3,4,5,
	}

	names = append(names, "coba")
	names = append(names, "coba2")
	names = append(names, "coba3")
	fmt.Println("--------ArrayDataType--------")

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(values)
	
}
