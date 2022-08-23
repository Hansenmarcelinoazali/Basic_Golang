package main

import (
	"container/list"
	"fmt"
)

/*
data yang saling terikat maupun sesudah atau sebelum data tersebut dan bentuknya array
*/

func main() {
	data := list.New()
	data.PushBack("hansen")
	data.PushBack("marcelino")
	data.PushBack("azali")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
