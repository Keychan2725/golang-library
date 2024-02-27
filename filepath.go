package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("Hello/Chandra.go"))
	fmt.Println(filepath.Base("Hello/Chandra.go"))
	fmt.Println(filepath.Ext("Hello/Chandra.go"))
	fmt.Println(filepath.IsAbs("Hello/Chandra.go"))
	fmt.Println(filepath.IsLocal("Hello/Chandra.go"))
	fmt.Println(filepath.Join("Hello", "Chan", "Chandra.go"))

}
