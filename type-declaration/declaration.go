package main

import "fmt"

//membuat alias pada suatu tipe data

func main(){
	type noKTP string
	var NoKTP noKTP = "12345678"
	fmt.Println(NoKTP)

	type married bool
	var marriedStatus married = true
	fmt.Println(marriedStatus)
}