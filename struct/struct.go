package main

import "fmt"

/*
1.struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data.
2. struct biasanya representasi dalam program aplikasi yang kita buat
3. data di struct di simpan dalam field
4. sederhananya struct adalah kumpulan dari field
*/
func main() {
	structs()
	StructLiteral()
}

type Customer struct {
	Name, Address string
	Age           int
}

func structs() {
	// var (
	// 	Name = "hansen"
	// 	Address = "yogya"
	// 	Age = 24
	// )
	// fmt.Println(Name,Address,Age)

	var hansen Customer

	hansen.Name = "hansen"
	hansen.Address = "yogya"
	hansen.Age = 24
	fmt.Println(hansen)
	fmt.Println(hansen.Name)
	fmt.Println(hansen.Address)
	fmt.Println(hansen.Age)
}

func StructLiteral() {

	fmt.Println("ini struct literal")
	stray := Customer{
		Name:    "cat",
		Address: "stray",
		Age:     24,
	}
	fmt.Println(stray)

	budi := Customer{
		"budi", "yogyakarta", 30,
	}
	fmt.Println(budi)
}
