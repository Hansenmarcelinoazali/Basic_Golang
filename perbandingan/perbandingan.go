package main

import "fmt"

func main() {
	var name1 = "hansen"
	var name2 = "hansen"

	var result bool = name1 == name2

	fmt.Println(result)

	var value1 int = 100
	var value2 int = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
