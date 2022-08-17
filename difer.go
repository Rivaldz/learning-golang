package main

import "fmt"

func logging()  {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication(value int)  {
	defer logging() //dieksekusi meskipun ada error
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("result = " , result)
}

func main()  {
	runApplication(10)
}