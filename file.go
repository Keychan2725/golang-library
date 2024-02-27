package main

import (
	"bufio"
	"io"
	"os"
)

func CreateNewFileName(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func addTofile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		Line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(Line) + "\n"
	}
	return message, nil

}
func main() {
	//CreateNewFileName("Wanderer app", "Hallo Player terima kasih telah memainkan gmae ini")

	//result, _ := readFile("Wanderer app")
	//fmt.Println(result)
	addTofile("Wanderer app", "\nini pesan tambahan")
}
