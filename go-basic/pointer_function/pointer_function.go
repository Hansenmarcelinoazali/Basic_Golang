package main

import "fmt"

/*
- Saat kita membuat parameter di function secara default adalah pass by value, artinya data akan di copy
lalu dikirim ke functio tersebut
- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah
- hal ini membuat variebl menjadi akan menjadi aman, karena tidak akan bisa diubah
- Untuk melakukan ini, kita juga bisa menggunakan pointer function
- Unutk menjadikan sebuah parameter sebagai function, kita bisa menggunakan operator * di parameternya
*/


type Address struct{
	City,Province, Country string
}

func ChangeAddresToIndonesia(address *Address){
	address.Country = "Indoensia"
}

func main()  {
	address :=Address{"Subang","Jawa Timur",""}

	ChangeAddresToIndonesia(&address)

	fmt.Println(address)
}