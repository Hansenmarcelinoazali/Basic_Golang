package main

import "fmt"

//function dengan mengembalikan data
func main()  {
	result:=getHello("hansen","marcelino azali")
	fmt.Println(result)
}

func getHello(name,lastname string)string{
	return "hallo "+name+lastname
}