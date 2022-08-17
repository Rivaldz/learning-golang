package main

import "fmt"

func getFullName2()(firstName string, midleName string, lastName string)  {
	firstName = "Muhammad"
	midleName = "Rivaldo"
	lastName = "Purnomo"

	return
}

func main()  {
 namaPertama, namaKedua, namaKeTiga := getFullName2()

 fmt.Println(namaPertama,namaKedua, namaKeTiga)

}