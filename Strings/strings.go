package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Hansen Marcelino","Hansen")) //cek value
	fmt.Println(strings.Split("Hansen Marcelino Azali"," ")) //memotong string berdasarkan spasi
	fmt.Println(strings.ToLower("Hansen Marcelino Azali")) //membuat menjadi lower
	fmt.Println(strings.ToUpper("hansen marcelinio azali")) //membuat menjadi kapital

	fmt.Println(strings.Trim("      hansen marcelino azali      "," "))
	fmt.Println(strings.ReplaceAll("eko eko eko eko eko", "eko","Budi"))
}