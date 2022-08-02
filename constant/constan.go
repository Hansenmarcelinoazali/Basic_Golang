package main

import "fmt"

//constant adalah variabel yang tidak bisa diubah nilainya

func main(){
	const first_name string = "hansen"
	const middle_name string = "marcelino"
	const last_name string = "azali"
	const umur int = 24
	
	fmt.Println(first_name,middle_name,last_name,"umur : ",umur)

//multiple constant
 const(
	satu = "Bank"
	dua = "Raya"
	tiga = "Indonesia"
 )
 fmt.Println(satu,dua,tiga)

}