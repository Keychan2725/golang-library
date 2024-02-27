package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Satria chandra pamungkas")
	data.PushBack(15)
	data.PushBack("Purbalingga")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	var Last *list.Element = head.Next()
	fmt.Println(Last.Value)

	next := Last.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)

	}
}
