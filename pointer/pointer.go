package main

import "fmt"

/*
Golang mengcopy value bukan by reference : jadi Variabel B akan mengcopy value variabel A
walaupun variabel B diubah nilai variabel A tidak akan berubah

*Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
sederhananya dengan kemampuan pointer, kita bisa membuat pass by reference

Operator &
- Untuk membuat sebauh variebl dengan nilai pointer ke variabel yang lain, kita bisa menggunakan operator & di ikuti dengan nama varieblnya

Operator *
-Saat kita mengubah variabel pointer, maka yang berubah hanya variabel tersebut
- Semua variebl yang mengacu ke data yang sama tidak akan berubah
- Jika kita ingin merngubah seluruh variebl yang mengacu ke data tersebut, kita bisa menggunakan operator *

Pointer New
-Mengambalikan pointer ke data kosong bukan data awal
*/

type Address struct {
	City, Province, Country string
}

func main() {
	//Pointer &
	address1 := Address{"Mojokerto", "Jambi", "Padang"}
	address2 := &address1
	address3 := &address1

	address2.City = "Yogyakarta" //mengubah pertama

	// address2 = &Address{"Jakarta","DKI Jakarta","Indonesia"} ////pointer&

	*address2 = Address{"Subang", "Jawa Timur", "Indonesia"} ////pointer *

	fmt.Println(address1) //tidak berubah karena by value bukan by reference
	fmt.Println(address2) //pointer
	fmt.Println(address3)

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Japan"
	// alamat2.City = "Tokyo"
	// alamat2.Province = "iuahs"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
