package files

import (
	"bufio"
	"fmt"
	"godesde0/exercises"
	"os"
)

var fileName = "./static/table.txt"

func RecordTable() {
	var text string = exercises.GenerateNumberBoard()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error to create file", err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func AddTable() {
	var text string = exercises.GenerateNumberBoard()
	if !Append(fileName, text) {
		fmt.Println(("error"))
	}
}

func Append(fileName string, text string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error to OpenFile", err.Error())
		return false
	}
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error WriteString", err.Error())
		return false
	}
	file.Close()
	return true
}

func ReadFile() {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error WriteString", err.Error())
		return
	}
	fmt.Println(string(file))
}

func ReadFile2() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error read file", err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		record := scanner.Text()
		fmt.Println("> ", record)
	}
	file.Close()
}
