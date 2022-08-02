package main

import "fmt"

func main() {

	//konversi tipe data integer
	var nilai32 int32 = 32758
	var nilai64 int64 = int64(nilai32) //mengubah nilai int32 ke int64

	var nilai16 int16 = int16(nilai32) //mengubah nilai dari int32 ke int16

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)


	fmt.Println("=============")
//konversi tipe data string
var name = "hansen"
var e = name[0] //saat mengambil dari suatu karakter akand ubah ke byte
var eString string = string(e)

fmt.Println(name)
fmt.Println(eString)
}