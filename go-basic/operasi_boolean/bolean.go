package main

import "fmt"

// && adalah dan
// || adalah atau
// ! adalah kebalikan

func main(){
var nilaiAkhir = 80
var absensi = 80

var lulusNilaiAkhir bool = nilaiAkhir >= 80
fmt.Println(lulusNilaiAkhir)

var lulusAbsensi bool = absensi >=80
fmt.Println(lulusAbsensi)

var lulusUjianAkhir bool = lulusNilaiAkhir && lulusAbsensi

fmt.Println(lulusUjianAkhir)
}