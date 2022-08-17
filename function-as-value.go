package main

import "fmt"

func getGoodBy(name string) string  {
	return "Good Byee " + name
}

func main()  {
	sayGoodBye := getGoodBy
	result := sayGoodBye("Rivaldo")
	fmt.Println(result)
	
}