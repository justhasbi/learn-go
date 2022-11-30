package main

import "fmt"

func Variable() {
	fmt.Println("Hello World")

	// golang data type
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)

	fmt.Println("False = ", false)
	fmt.Println("True = ", true)

	fmt.Println("Hasbi")
	fmt.Println("Hasbi Shuhada")
	fmt.Println(len("hasbishuhada"))
	fmt.Println("hasbi"[0])

	// golang variable
	var name string = "hasbi"
	name = "hasbi shuhada"
	fmt.Println(name)

	firstName := "shuhada"
	fmt.Println(firstName)

	firstName = "gunawan"
	fmt.Println(firstName)
	// declare multiple variable
	var (
		age     = 20
		address = "tegal"
	)

	fmt.Println(age)
	fmt.Println(address)

	// Constant
	const (
		APPLICATIONNAME = "Learn GO"
		APPLICATION     = "GOLANG"
	)
	// cannot assign to constant
	//APPLICATIONNAME = "coba"

	fmt.Println(APPLICATIONNAME)
	fmt.Println(APPLICATION)
}
