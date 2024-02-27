package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvstring := "Chandra,Pamungkas.Wanderer\n" +
		"Novian,Chandra,Chan\n" +
		"Wanderer,Chan,Chandra"

	reader := csv.NewReader(strings.NewReader(csvstring))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else {
			fmt.Print(record)
		}
	}
}
