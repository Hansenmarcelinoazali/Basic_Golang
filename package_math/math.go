package main

import (
	"fmt"
	"math"
)

func main()  {

	// jika diatas 0.5 maka akan membualatkan keatas
	// jika dibawah 0.5 maka akan membulatkan ke bawah
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	//floor akan memaksa membulatkan kebawah
	fmt.Println(math.Floor(1.7))
	//ceil akan memaksa membulatkan keatas
	fmt.Println(math.Ceil(1.3))

	// max dan min
	fmt.Println(math.Max(10,20))
	fmt.Println(math.Min(10,20))
}