package main

import "fmt"

func main(){
// sayHello()
test()

}

func sayHello(){
	fmt.Println("hello world")
}

func test(){
	for i:=0;i<=10;i++{
		sayHello()
	}
}