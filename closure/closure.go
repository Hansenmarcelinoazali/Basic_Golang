package main

import "fmt"


//closure func adalah fungsi yang bisa berinteraksi dengan value yang ada di atasnya
func main()  {
	counter :=0

	increment:= func()  {
		fmt.Println("increment")
	}
	increment()
	fmt.Println(counter)
}