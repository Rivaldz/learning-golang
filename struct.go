package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var aldo Customer
	aldo.Name = "Rivaldo"
	aldo.Address = "Dusun Bogo"
	aldo.Age = 23

	fmt.Println(aldo)
	//print tiap field
	fmt.Println(aldo.Name)
	fmt.Println(aldo.Address)
	fmt.Println(aldo.Age)

	purnomo := Customer{
		Name: "Muhammad",
		Address: "Bogo",
		Age: 23,
	}

	fmt.Println(purnomo)

	setyo := Customer{"Aldo","Bogo",23}
	fmt.Println(setyo)

}
