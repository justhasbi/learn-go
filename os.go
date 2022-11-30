package main

import (
	"fmt"
	"os"
)

func osPackage() {
	args := os.Args
	fmt.Println("Argument: ")
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])
	name, error := os.Hostname()
	if error == nil {
		fmt.Println(name)
	} else {
		fmt.Println(error.Error())
	}
}
