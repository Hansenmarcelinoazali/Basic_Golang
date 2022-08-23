package main

import "fmt"

/*
type assertion itu mengubah tipe data ke tipe data yang diinginkan
fitur ini serin gsekali digunakan ketika kita bertemu dengan data interface kosong
*/

func Random() interface{} {
	return true
}

func main() {
	//untuk melakukan tipe assertion harus menentukan juga tipe data interface kosong tersebut
	result := Random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)//panic
	// fmt.Println(resultInt)
//////////
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("integer", value)
	default:
		fmt.Println("unknow type")
	}

	
}
