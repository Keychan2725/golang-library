package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation erorrs")
	NotfoundError   = errors.New("Notfound errors")
)

func GetBYId(id string) error {

	if id != "chandra" {
		return NotfoundError
	}
	if id == "" {
		return ValidationError
	}
	return nil
}
func main() {
	err := GetBYId("chandra")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation erorrs")
		} else if errors.Is(err, NotfoundError) {
			fmt.Println("Notfound error")
		} else {
			fmt.Println("Unkwon errors")
		}

	}

}
