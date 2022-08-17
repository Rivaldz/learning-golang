package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	//akan terjadi overflow dan di set nilai paling rendah
	fmt.Println(nilai8)

	var name = "aldoo"
	//ini adalah uint/byte
	var a = name[0]
	var conva string = string(a)
	fmt.Println(conva)
}
