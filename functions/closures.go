package functions

import "fmt"

func Table(value int) func() int {
	number := value
	sequence := 0
	return func() int {
		sequence++
		return number * sequence
	}
}

func CallClosure() {
	myTable := Table(2)
	for i := 1; i <= 10; i++ {
		fmt.Println(myTable())
	}
}
