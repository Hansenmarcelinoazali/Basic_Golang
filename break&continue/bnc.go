package main

import "fmt"

//break untuk memberhentikan perulangan
//continue untuk skip
func main(){
// for i:=0;i <10;i++{
// 	fmt.Println("perulangan i:",i)
// 	if i ==5{
// 		break
// 	}
// }
for i:=0; i<10;i++{
	if i%2 == 0{
		continue
	}
	fmt.Println(i)
}

}