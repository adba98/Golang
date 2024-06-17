package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var caption string
var err error

func InputNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter number 1:")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Incorrect data" + err.Error())
		}
	}

	fmt.Println("Enter number 2:")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Incorrect data" + err.Error())
		}
	}

	fmt.Println("Enter caption:")
	if scanner.Scan() {
		caption = scanner.Text()
		if err != nil {
			panic("Incorrect data" + err.Error())
		}
	}

	fmt.Println(caption, number1*number2)
}
