package main

import "fmt"

//defer
func logging(){
	fmt.Println("selesai memanggil function")
}

func runApplication(){
	defer logging()
	fmt.Println("run application")
}

func aplikasi(value int){
	defer logging()
	result := 10/value
	fmt.Println("result",result)
	
}

//panic untuk menghentikan program jika terpanggil
//tapi defer tetap di eksekusi
func endapp(){
	// fmt.Println("end app")
	message := recover()
	fmt.Println("terjadi error",message)
}

func Panic(error bool){
	defer endapp()
	if error == true{
		panic("ERROR")
	}
}

//recover digunakan untuk menangkap data panic
//dengan recover program akan tetap berjalan
func appRecover(error bool){
	defer endapp()
	if error == true{
		panic("error")
	}
	
	fmt.Println("aplikasi berjalan")
}

func main(){
	runApplication()

	aplikasi(10)

	Panic(false)

	appRecover(true)
}