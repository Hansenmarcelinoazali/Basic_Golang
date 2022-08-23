package main

import (
	"fmt"
	"sort"
)


type User struct{
	Name string
	Age int
}

type UserSlice []User

func (value UserSlice)Len()int{
	return len(value)
}

func(Value UserSlice)Less(i,j int)bool{
	return Value[i].Age < Value[j].Age
}

func (userSlice UserSlice) Swap(i,j int) {
	userSlice[i],userSlice[j] = userSlice[j],userSlice[i]
}

func main(){
users := []User{
	{"hansen",24},
	{"hendro",24},
	{"aldi",30},
}

sort.Sort(UserSlice(users))

fmt.Println(users)
}