package main

import (
	"bufio"
	"os"
)

func main() {
	input := bufio.NewWriter(os.Stdout)
	_, _ = input.WriteString("Hello Wanderer\n")
	_, _ = input.WriteString("I love You \n")
	input.Flush()

}
