package main

import (
	"fmt"
	"strconv"
)

//Konversi tipe data

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
	val := strconv.FormatInt(1000000,2)
	fmt.Println(val)

	inttostr := strconv.Itoa(1000000) //int to string
	fmt.Println("ini inttostr",inttostr)

	strtoint,_ := strconv.Atoi("1000000")
	fmt.Println("ini strtoint",strtoint)


}
