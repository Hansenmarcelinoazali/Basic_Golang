package main

import "fmt"

// array adalah tipe data dengan kumpulan tipe data yang sama
// array nilainya sudah di tetapkan
func main() {
	var namaLengkap [3]string
	namaLengkap[0] = "hansen"
	namaLengkap[1] = "marcelino"
	namaLengkap[2] = "azali"

	fmt.Println(namaLengkap[0])
	fmt.Println(namaLengkap[1])
	fmt.Println(namaLengkap[2])
	fmt.Println("ini adalah panjang data dari nama lengkap : ",len(namaLengkap))

	//array int
	var values = [3]int{
		1,2,3,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println("ini adalah panjang values : ",len(values))
}
