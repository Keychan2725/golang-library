package main

import (
	"fmt"
	"strconv"
)

func main() {

	parsebool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(parsebool)
	}

	result, err := strconv.Atoi("213123")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(result)
	}
	format := strconv.FormatInt(999, 2)
	fmt.Println(format)
}
