package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MySlowName(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
