package main

import (
	"fmt"
)

func main() {
	var name string = "iwan"

	switch name {
	case "hansen":
		fmt.Println("hallo hansen")
	case "puji":
		fmt.Println("hallo puji")
	case "hendro":
		fmt.Println("hallo hendro")
	default:
		fmt.Println("hallo, kamu siapa?")
	}

	fmt.Println("======================")

	//dengan kondisi
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama terlalu pendek")
	}

	//tanpa kondisi
	len := len(name)

	switch {
	case len > 10:
		fmt.Println("nama terlalu panjang")
	case len > 5:
		fmt.Println("nama cukup panjang")
		default :
		fmt.Println("nama sudah pas")
	}
}
