package main

import "fmt"

func MapDataType() {
	var person map[string]string = map[string]string{
		"name":    "Hasbi",
		"address": "Tegal",
	}

	person["email"] = "hasbi@mail.com"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["email"])

	// make function
	fmt.Println("===============")
	book := make(map[string]string)

	book["title"] = "harry potter"
	book["author"] = "JK Rowling"
	book["wrong"] = "wrong"

	fmt.Println(book["title"])
	fmt.Println(book["author"])
	fmt.Println(book["wrong"])

	delete(book, "wrong")
	fmt.Println(book)
}
