package main

import (
	"fmt"
	// use blank identifier _ if you just want to implement
	// initialization func in package
	"learn-go/database"
)

func packageInit() {
	result := database.GetDatabase()

	fmt.Println(result)
}
