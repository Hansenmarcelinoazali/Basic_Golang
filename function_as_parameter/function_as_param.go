package main

import (
	"fmt"
)

func main() {
	//1
	sayHelloWithFilter("hansen", spamFilter)

	//2
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("hello ", filter(name))
}

func spamFilter(names string) string {
	if names == "Anjing" {
		return "..."
	} else {
		return names
	}
}
