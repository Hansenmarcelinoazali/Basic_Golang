package main

import "fmt"

//named return value

func getCompleteNames() (firstname string, middlename string, lastname string) {
	firstname = "hansen"
	middlename = "marcelino"
	lastname = "azali"

	return
}

func main() {
	firstname, middlename, lastnames := getCompleteNames()
	fmt.Println(firstname, middlename, lastnames)
}
