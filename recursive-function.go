package main

import "fmt"

func factorialRecursif(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursif(value-1)
	}
}

func main() {
	factorialRecursif := factorialRecursif(10)
	fmt.Println(factorialRecursif)
}
