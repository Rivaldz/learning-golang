package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Rivaldo"
	names[2] = "Setyo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values =  [4]int{
		4,
		5,
		6,
		5,
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

}
