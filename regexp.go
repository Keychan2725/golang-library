package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile(`c([a-s])a`)

	fmt.Println(regex.MatchString("chantika"))
	fmt.Println(regex.MatchString("chandira"))
	fmt.Println(regex.MatchString("ckana"))

}
