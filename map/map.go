package main

import "fmt"

//map adalah tipe kumpulan data lain yang berisi data yang tipe data sama

func main()  {
	//untuk membuat map harus mendifinisikan tipe data key dan tipe data valuenya
	person:=map[string]string{ 
		"name": "hansen",
		"address" : "Yogyakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "programmer"
	fmt.Println()
}