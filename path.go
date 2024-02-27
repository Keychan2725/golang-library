package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("Hello/Hello.go"))
	fmt.Println(path.Base("Hello/Hello.go"))
	fmt.Println(path.Ext("Hello/Hello.go"))
	fmt.Println(path.Join("Hello", "Chan", "Wanderer.go"))

}
