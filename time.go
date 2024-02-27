package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2008, time.May, 27, 16, 30, 27, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2008-05-27T15:04:05Z")
	fmt.Println(parse)
}
