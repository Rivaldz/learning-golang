package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice = months[4:7]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//merubah array dari slace
	// slice[0] = "uwowww"
	// fmt.Println(months)

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Mas Aldo")
	fmt.Println(slice3)
	//mengubah data di dalam
	slice3[1] = "Data ubahan"
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Muhamad Rivaldo"
	newSlice[1] = "Muhamad Rivaldo Setyo"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
