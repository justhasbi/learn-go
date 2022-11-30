package main

import (
	"fmt"
	"learn-go/Helper"
)

func other() {
	fmt.Println("Importing helper...")
	Helper.CobaImport()
}

func otherfunc() {
	other()
}
