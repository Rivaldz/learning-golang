package main

import "fmt"

func main() {
	const firstName string = "Muhammad"
	const lastName = "Rivaldo"

	//error tidak bisa diubah
	// firstName = "uwuwu"
	// lastName = "cuakss"

	//Tidak error ketika tidak digunakan
	fmt.Println(firstName)
	fmt.Println(lastName)

	//multiple const
	const (
		multione = "one"
		multitwo = "two"
	)
	fmt.Println(multione)
	fmt.Println(multitwo)

}
