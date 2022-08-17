package main

import "fmt"

func main()  {
name := "Rivaldo"
counter := 0

increment := func ()  {
	name := "Budi" //tidak akan mengubah yang bagian atas
	fmt.Println("Increment")
	counter++ //akan mengubah bagian yang atas
	fmt.Println(name)
}
increment()
increment()

fmt.Println(name)	
fmt.Println(counter)
}