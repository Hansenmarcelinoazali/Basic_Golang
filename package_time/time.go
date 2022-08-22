package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())

	layout := time.RFC3339
	parsing, err :=time.Parse(layout,"2022-08-22")
	if err == nil{
		fmt.Println(parsing)
	}else{
		fmt.Println(err.Error())
	}
}