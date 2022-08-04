package main

import "fmt"

func main() {
	// Possmov(Position{1, 7}) //Something
	// values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// x:= tambah(values)
	// fmt.Println(x)
	gg()

}

// type Position struct {
//     row int
//     col int
// }

// func (posstn Position) isvalid() bool {
//     if posstn.row > 8 || posstn.row < 0 || posstn.col > 8 || posstn.col < 0 {
//         return false
//     }
//     return true
// }

// func Possmov(pos Position) {
//     isval := pos.isvalid()
//     if isval == true {
//         fmt.Println("Something")
//     }
// }

// func tambah() (result int) {
// 	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	// result := 0
// 	for i := range values {
// 		result += i
// 	}
// 	fmt.Println(result)
// 	return result
// }

func calc() (int, int) {
	return 1, 2
}

func gg() {
	a, _ := calc()
	fmt.Println(a)
}
