package main

import "fmt"

func getFullName()(string,string,string)  {
	return "Muhammad", "Rivaldo", "Setyo"
	
}

func main()  {
	firsName,_,lastName := getFullName()

	fmt.Println(firsName,lastName)

	
}