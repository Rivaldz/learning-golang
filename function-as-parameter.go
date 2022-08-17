package main

import "fmt"

// func sayHelloWithFiler(name string, filter func(string) string)  {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello " + nameFiltered)
// }

//Bisa menggunakan alias func untuk menangani data constructor yang panjang
type Filter func(string)string

func sayHelloWithFiler(name string, filter Filter)  {
	nameFiltered := filter(name)
	fmt.Println("Hello " + nameFiltered)
}

func spamFilter(name string) string{
	if  name == "Asu" {
		return "----"	
	}else{
		return name
	}
}

func main()  {
	sayHelloWithFiler("Rivaldo", spamFilter)
	sayHelloWithFiler("Asu", spamFilter)
	
}