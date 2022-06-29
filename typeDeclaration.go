package main

import "fmt"

func TypeDeclaration() {
	type KTPNumber string
	type MaritalStatus bool

	var noKTPHasbi KTPNumber = "1233111331311"
	var hasbiMaritalStatus MaritalStatus = false
	fmt.Println("--------TypeDeclaration--------")

	fmt.Println(noKTPHasbi)
	fmt.Println(hasbiMaritalStatus)
}
