package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Rivaldo"
	fmt.Println(name)

	name = "Setyo Purnomo"
	fmt.Println(name)

	//Bisa juga menggunakan variable tanpa tipe data
	var nama = "Setyo Purnomo nama"
	fmt.Println(nama)

	nama = "Aldo nama"
	fmt.Println(nama)

	//Kata kunci var tidak wajib
	namaTeman := "Setyo Purnomo Teman"
	fmt.Println(namaTeman)

	namaTeman = "Aldo Teman"
	fmt.Println(namaTeman)

	//Multi Variable
	var(
		firstName = "first name = Muhammad"
		lastName = "last name = Purnomo"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}
