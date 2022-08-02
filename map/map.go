package main

import "fmt"

//map adalah tipe kumpulan data lain yang berisi data yang tipe data sama

func main() {
	//untuk membuat map harus mendifinisikan tipe data key dan tipe data valuenya
	person := map[string]string{
		"name":    "hansen",
		"address": "Yogyakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "programmer"
	fmt.Println("ini adalah person : ", person)

	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	fmt.Println("ini adalah chiken : ", chickens)

	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	fmt.Println("ini adalah data:", data)
}
