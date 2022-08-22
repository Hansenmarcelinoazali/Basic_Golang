package main

import (
	"fmt"
	"os"
)

/*
- Golang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
- Package OS berisikan funsional untuk mengakses fitur sistem operasi secara independen bisa digunakan di semua sistem operasi
*/

func main(){
	args := os.Args
	fmt.Println("Argumen: ", args)

	hostname, err :=os.Hostname()
	if err == nil{
		fmt.Println("ini adalah hostname",hostname)
	}else{
		fmt.Println("ini adalah error",err.Error())
	}
}