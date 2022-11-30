package main

import "fmt"

func SliceDataType() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[6:8]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniArray2 := [...]int{1, 2, 3, 4, 5, 6}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)

}
