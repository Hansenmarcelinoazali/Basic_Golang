package main

import "fmt"

/*
1. Biasanya di dalam bahasa pemograman lain, object yang belum diinisialisaki maka secara otomatis nilainya adalah null atau nill
2. Berbeda dengan Go-Lang, di golang saat kita buat variabel dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya
3. Namun di Golang ada data nil, yaitu data kosong
4. Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
*/

func NewMap(name string) map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"name":name,
		}
	}
}

func main(){
	// person := NewMap("hansen")
	// fmt.Println(person)

	fmt.Println(NewMap("hansen"))
}