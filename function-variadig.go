package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func main() {
	total := sumAll(5,5,5,5,5)
	fmt.Println(total)

	numbers := []int{1,2,3,4,5}
	total = sumAll(numbers...)
	fmt.Println(total)
}
