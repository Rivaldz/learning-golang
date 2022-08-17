package main

import "fmt"

func main(){

	type NoKTP string
	type Married bool

	var noKTPAldo NoKTP = "293098409283908492830498"
	var status Married = false

	fmt.Println(noKTPAldo)
	fmt.Println(status)

}