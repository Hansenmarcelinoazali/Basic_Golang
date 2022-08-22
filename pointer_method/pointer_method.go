package main

import "fmt"

/*
Pointer di Method
- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
- Sangat direkomendasikan menggnakan poiunter di method, sehingga tidak boros memory karena harus di duplukasi ketika memanggil method
*/

type Man struct{
	Name string
}

func(man *Man) Married(){
	man.Name = "Mr. "+man.Name
	fmt.Println("di method",man.Name)
}

func main(){
	hansen:= Man{"Hansen"}
	hansen.Married()
	fmt.Println(hansen.Name)
}