package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("Aplikasi Selesai")
	/* recover function digunakan untuk menangkap data panic
	dengan function ini maka panik akan terhenti dan program
	tetap akan dijalakan
	*/
	message := recover() // menangkap error panic
	if message != nil {
		fmt.Println("terjadi ERROR: ", message)
	}
}

func runApp(error bool) {
	/* defer digunakan agar setiap selesai execute function maka akan
	otomatis menjalankan function defer
	*/
	defer endApp()
	if error {
		/* panic function yang dapat digunakan untuk menghentikan program
		funtion ini biasanya dipanggil ketika terdapat error pada saat program berjalan
		ketika panic dipanggil, defer function tetap akan dijalankan
		*/
		panic("Error")
	}
	fmt.Println("Selesai")
}

func DeferPanicRecover() {
	runApp(true)
}
