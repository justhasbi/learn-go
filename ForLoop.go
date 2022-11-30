package main

import "fmt"

func ForLoop() {
	for counter := 1; counter <= 15; counter++ {
		fmt.Println(counter)
	}

	slice := []string{"hasbi", "laeli", "abdul"}

	for i := 0; i < len(slice); i++ {

	}

	for index, item := range slice {
		fmt.Println("index:", index, "=", item)
	}
}
