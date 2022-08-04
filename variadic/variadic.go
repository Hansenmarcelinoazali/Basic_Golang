package main

import (
	"fmt"
)

func sumAll(numbers ...int) int {
	total := 0
	for _, values := range numbers {
		total += values
	}
	return total
}

func main() {
	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := sumAll(number...)
	fmt.Println(result)
}
