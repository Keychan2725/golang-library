package main

import (
	"encoding/csv"
	"os"
)

func main() {
	csvWriter := csv.NewWriter(os.Stdout)

	_ = csvWriter.Write([]string{"Chan", "Wanderer", "Pamungkas"})
	_ = csvWriter.Write([]string{"Chan", "Yaemiko", "Key"})
	_ = csvWriter.Write([]string{"Keychan", "Chandra", "CHANDER"})

	csvWriter.Flush()

}
