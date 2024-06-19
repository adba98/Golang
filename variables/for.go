package variables

import "fmt"

func Iterate() {
	for i := 0; i < 50; i += 5 {
		if i == 25 {
			continue
		}
		fmt.Println(i)
	}
}
