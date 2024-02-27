package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("satria", "sat"))
	fmt.Println(strings.Split("satria chandra", " "))
	fmt.Println(strings.ToLower("satria chandra"))
	fmt.Println(strings.ToUpper("satria chandra"))
	fmt.Println(strings.Trim("        satria chandra      ", " "))
	fmt.Println(strings.ReplaceAll("pamungkas satria chandra pamungkas", "pamungkas", "Keychan"))

}
