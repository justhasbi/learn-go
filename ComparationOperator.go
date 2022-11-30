package main

import "fmt"

func ComparationOperation() {
	var name1 = "hasbi"
	var name2 = "hasbi"

	var (
		val1    = 1
		val2    = 3
		nilai1  = 81
		absensi = 81
	)

	var result bool = name1 == name2
	fmt.Println("--------ComparationOperation--------")
	fmt.Println(result)
	fmt.Println(val1 > val2)
	fmt.Println(val1 >= val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 <= val2)
	fmt.Println(val1 != val2)

	fmt.Println("--------BooleanOperations--------")
	if nilai1 > 80 && absensi > 80 {
		fmt.Println("lulus")
	} else {
		fmt.Println("gagal")
	}
}
