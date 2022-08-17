package main

import "fmt"

func main() {
	name := "Muhamad Rivaldo"

	switch name {
	case "Aldo":
		fmt.Println("Haloo Doo")
	case "Rivaldo":
		fmt.Println("Haloo RiDoo")
	default:
		fmt.Println("Mbuhlahh")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Sudah Benar")
	case false:
		fmt.Println("Nama Salah")

	}

	lengthName := len(name)
	switch {
	case lengthName < 6:
		fmt.Println("Nama Pendek")
	case lengthName > 6:
		fmt.Println("Nama Panjang")
	case lengthName < 1:
		fmt.Println("Nama Nggak Niat")
	}
}
