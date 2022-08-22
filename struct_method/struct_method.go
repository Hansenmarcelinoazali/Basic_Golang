package main

import "fmt"

/*
Struct Method
1. Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagau parameter untuk function
2. namun juka kita ingin menambahkan method ke dalam struct, shingga seakan-akan sebuah struct memiliki function
3. method adalah function
*/
type Customer struct{
	Names, Address string
	Age int
}

func(customer Customer)sayHi(name string){
	fmt.Println("hello",name, "my name is",customer.Names)
}

func main()  {
var customer Customer 
	customer.Names = "hansen"
	customer.Address = "yogya"
	customer.Age = 24

customer.sayHi("BOT")

}