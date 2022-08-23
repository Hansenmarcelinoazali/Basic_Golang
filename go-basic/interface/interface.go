package main

import "fmt"

/*
1. Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
2. Sebuah interface berisikan defininsi-definisi method
3. Biasanya interface digunakan sebagai kontrak
*/

type HasName interface {
	GetName() string
}

func sayHello(hasnames HasName) {
	fmt.Println("hello", hasnames.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var hansen Person
	hansen.Name = "hansen"

	sayHello(hansen)
}
