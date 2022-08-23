package main

import "fmt"

//recrusive func adalah func yang memanggil dirinya sendiri
func main() {
	fmt.Println("ini loop:",factorialLoop(10))
	fmt.Println("ini recursive:",factorialRecursive(10))
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int)int{
	if value == 1{
		return 1
	}else{
		return value *factorialRecursive(value-1)
	}
}
