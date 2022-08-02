package main

import (
	"fmt"
)

// tipe data slice adalah potongan dari array
// slice mirip dengan array yang membedakan adalah jumlah dari datanya
//slice dan array sama cara mengambill nilai dari datanya

//detail tipe data slice
//1. tipe data slice memiliki 3 data, yaitu pointer, length dan capacity
//2. pointer adalah penunjuk data pertama di array para slice
//3 .length adalah panjang dari slice dan capacity adalah kapasitas dari slice dimana length tidak boleh lebih dari capacity

func main() {
	// var values = []int{1,2,3,4,4,5,5}

	month := [...]string{
		"januari",   //0
		"februari",  //1
		"maret",     //2
		"april",     //3
		"mei",       //4
		"juni",      //5
		"juli",      //6
		"agustus",   //7
		"september", //8
		"oktober",   //9
		"november",  //10
		"desember"}  //11
	slice1 := month[4:7]
	fmt.Println(slice1)                                //mengambil nilai 4 5 6 yaitu mei sampai juli
	fmt.Println("ini adalah len slice1", len(slice1))  //valuenya ada 2
	fmt.Println("ini adalah cap slice 1", cap(slice1)) //kapasitas dari 4 sampai akhir

	month[5] = "diubah"
	fmt.Println(slice1)

	slice1[0] = "mei lagi"
	fmt.Println(month)

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu"}
	daySlice := days[5:] //mengamil nilai dari ke 5
	fmt.Println(daySlice)

	daySlice[0] = "sabtu baru"  //mengganti value index 0
	daySlice[1] = "minggu baru" //mengganti value index 1
	fmt.Println(days)

	daySlice2 := append(daySlice, "libur baru") //menambah value dari index trakhir
	daySlice2[0] = "upps"                       //menambah value pada index 0

	fmt.Println(daySlice2)
	fmt.Println(days)

	//membuat slice baru
	fmt.Println("=== Slice Baru ===")
	newSlice := make([]string, 2, 5) //2 adalah length//5 adalah capacity
	newSlice[0] = "hansen"
	newSlice[1] = "marcelino"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string,len(newSlice), cap(newSlice)) //length dan capacitynya harus dipastikan sama kalau tidak nanti terpotong
	copy(copySlice, newSlice)
	fmt.Println("ini adalah copy slice: ",copySlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println("ini adalah array: ",iniArray)
	fmt.Println("ini adalah slice: ",iniSlice)

	//array mendefinisikan jumlah data yang ditampung
}
