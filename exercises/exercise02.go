package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GenerateNumberBoard() {
	var number int
	var err error
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
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}
