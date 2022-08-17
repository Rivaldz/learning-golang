package main

import "fmt"

func main()  {
	name := "Al"

	if name == "Al" {
		fmt.Println("Hello Al")
	}else if name == "Do"{
		fmt.Println("Hello Do")
	}else if name == "DI"{
		fmt.Println("Hello DI")
	}else{
		fmt.Println("Woyyyy")
	}

	//short statement
	if length := len(name) ; length > 5{
		fmt.Println("Nama Terlalu Panjang")
		fmt.Println(length)
	}else{
		fmt.Println("Nama OKE")
	}
	
	
}