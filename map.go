package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":   "Rivaldo",
		"Person": "Nganjuk",
	}

	person["tile"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["Person"])

	var book map[string]string = make(map[string]string)

	book["title"] = "Woles"
	book["tahun"] = "2000"
	book["ups"] = "whahahah"

	delete(book,"ups")

	fmt.Println(book)

}
