package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan Ke " , counter)
	// 	counter++
	// }

	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan Ke ", i)
	}

	slice := []string{"Muhammad", "Rivaldo", "Setyo", "Purnomo"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	// bisa menggunakan _ untuk variable yang tidak digunakan
	for i, value := range slice {
		fmt.Println("Index", i, "Value", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)

	person["neme"] = "Rivaldo"
	person["age"] = "23"
	person["born"] = "99"

	for key, value := range person{
		fmt.Println(key , "Value", value)
	}

}
