package main

import (
	"fmt"
	"math"
)

func main(){
firstname,_:=getFullName()
fmt.Println("hallo ",firstname)


fmt.Println("==========")

var diamater float32 =15

area,circumference:=calculate(float64(diamater))
fmt.Println("luas lingkaran ", area)
fmt.Println("keliling lingkaran ",circumference)
}

func getFullName()(string, string){
	return "hansen","marcelino"
}

func calculate(d float64) (float64,float64){
	var area = math.Pi *math.Pow(d /2,2)

	var circumference =math.Pi*d

	return area,circumference
}
