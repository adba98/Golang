package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GenerateNumberBoard() string {
	var number int
	var err error
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter number:")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			break
		}
	}
	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d\n", number, i, number*i)
	}
	return text
}
