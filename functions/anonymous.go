package functions

import "fmt"

func Calculations() {
	var num1 = 5
	var num2 = 10
	calculation := func(num1 int, num2 int) int {
		return num1 + num2
	}

	fmt.Println(calculation(num1, num2))

	calculation = func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Println(calculation(num1, num2))
}
