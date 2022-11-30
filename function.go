package main

import "fmt"

func getFullname() (string, string, string) {
	return "hasbi", "shuhada", "best"
}

func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "hasbi"
	middlename = "shuhada"
	lastname = "shuhada"
	return
}

func sumAll(numbers ...int) int {
	total := 0
	for _, i := range numbers {
		total += i
	}

	return total
}

func getGoodbye(name string) string {
	return "Goodbye " + name
}

// function as parameter
type Filter func(string)string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("helo " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

// anonymous function
type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}


func function() {
	_, lastname, best := getFullname()
	fmt.Println(lastname, best)

	_, middlename, _ := getCompleteName()
	fmt.Println(middlename)

	sumTotal := sumAll(10, 10, 20, 30, 40)
	fmt.Println(sumTotal)

	slice := []int{12, 3, 4, 5, 6, 7}
	fmt.Println(sumAll(slice...))

	// save function in variable
	goodbye := getGoodbye
	fmt.Println(goodbye("hasbi"))

	// function as paramter
	filter := spamFilter
	sayHelloWithFilter("anjing", filter)

	// anonymous function
	blacklist := func(name string) bool  {
		return name == "anjing"
	}

	registerUser("Hasbi", blacklist)
	registerUser("Anjing", func(s string) bool {
		return s == "Anjing"
	})
}
