package database

import "fmt"


var connection string

func init(){
	connection = "MySQL berhasil di panggil"
	fmt.Println("berhasil di panggil")
}

func GetDatabase()string  {
	return connection
}