package main

import (
	"fmt"
)

//anonymous adalah function tanpa nama func
func main() {
	blcklst := func(name string) bool {
		return name == "anjing"
	}

	registerUser("hansen", blcklst)
	
	// registerUser("anjing", func(name string) bool {
	// 	return name == "anjing"
	// })


}

type Blacklist func(string) bool

// func blacklistAdmin(name string) bool{
// 	return name =="Admin"
// }
// func blacklistRoot(name string)bool{
// 	return name =="Root"
// }

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are block: ", name)
	} else {
		fmt.Println("welcome: ", name)
	}
}
