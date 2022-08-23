package main

import (
	"errors"
	"fmt"
)

/*
golang punya interface error
*/


func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian nilai 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main(){
	hasil, err := Pembagian(100,10)
	if err == nil{
		fmt.Println("hasil", hasil)
	}else{
		fmt.Println(err.Error())
	}

}