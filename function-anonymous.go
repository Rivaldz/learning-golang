package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Welocome", name)
	} else {
		fmt.Println("You are blocked")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Admin"
	}

	registerUser("Admin", blacklist)
	registerUser("Rivaldo", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("main", func(name string) bool {
		return name == "root"
	})
}
