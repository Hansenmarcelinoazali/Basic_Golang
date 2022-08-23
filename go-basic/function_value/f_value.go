package main

import "fmt"

//function as variabel

func main()  {
getgoodbye := getGoodBye
fmt.Println(getgoodbye("hansen"))

ambil()
	
}

func getGoodBye(name string) string{
	return "GoodBye " + name

}

func ambil(){
	ambil:=getGoodBye
	fmt.Println(ambil("iwan"))
}