package main

import "fmt"

func getHello(name string)string  {
	if name == "Rivaldo" {
		return "Hallo mr " + name
	}else{
		return "Haloo broo " + name
	}
}
func main()  {
	nameBro := getHello("Diooo")
	fmt.Println(getHello("Rivaldo"))
	fmt.Println(nameBro)
	
}