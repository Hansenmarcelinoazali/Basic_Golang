package main

import "fmt"

func main(){
	name:="iwan"

	if name == "hansen"{
		fmt.Println("hallo hansen")
	}else if name == "puji" {
		fmt.Print("hallo puji")
	}else{
		fmt.Println("kamu siapa?")
	}

	var length = len(name)
	if length > 5{
		fmt.Println("abc")
	}else if length < 5 {
		fmt.Println("def")
	}
}