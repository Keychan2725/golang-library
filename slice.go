package main

import (
	"fmt"
	"slices"
)

func main() {
	nama := []string{"Satria chandra", "Pamungkas", "Wanderer"}
	value := []int{10, 324, 555, 0}

	fmt.Println(slices.Max(value))
	fmt.Println(slices.Min(value))
	fmt.Println(slices.Contains(nama, "Wanderer"))
	fmt.Println(slices.Index(nama, "Pamungkas"))

}
