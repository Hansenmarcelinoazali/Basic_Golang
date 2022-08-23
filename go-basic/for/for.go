package main

import (
	"fmt"
)

func main() {

	//ini adalah for biasa
	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("ini adalah counter for 2", counter)
	}

	//ini adalah for-range

	var slice = []string{"hansen", "marcelino", "azali"}

	for i := 0; i < len(slice); i++ {
		fmt.Println("ini adalah i:", slice[i])
	}
	fmt.Println("==========")
	for index, value := range slice {
		fmt.Println("ini adalah for range value", value)
		fmt.Println("ini adalah index for range", index)
	}
fmt.Println("==============")
	person := make(map[string]string)
	person["name"] = "hansen"
	person["umur"] = "24 tahun"
	for r, v := range person {
		fmt.Println("index : ", r, "values : ", v)
	}

	fmt.Println("==================")
	orang := map[string]string{
		"nama": "hansen", "umur": "24 tahun",
	}
	for x, y := range orang {
		fmt.Println("index", x, "values:", y)
	}
}
