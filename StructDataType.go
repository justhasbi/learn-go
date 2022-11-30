// struct bisa dibilang sebagai class dalam golang
// yang didalamnya merupakan kumpulan field
package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

/** struct method
penggunaan struct method ini disarankan ketika kita ingin membuat method
yang spesifik untuk struct tertentu
*/
func (customer Customer) greeting() {
	fmt.Println("Hello", customer.Name)
}

func StructDataType() {
	// var hasbi Customer
	// hasbi.Address = "Tegal"
	// hasbi.Name = "Hasbi Shuhada"
	// hasbi.Age = 23

	// fmt.Println(hasbi)
	// fmt.Println(hasbi.Address)
	// fmt.Println(hasbi.Name)
	// fmt.Println(hasbi.Age)

	// joko := Customer {
	// 	Name: "Joko Pinurbo",
	// 	Address: "Palangkaraya",
	// 	Age: 45,
	// }
	// fmt.Println(joko)

	// budi := Customer { "Budi Doremi", "Jakarta", 32 }
	// fmt.Println(budi)

	gina := Customer{"Gina ayu", "Jakarta", 25}

	gina.greeting()
}
