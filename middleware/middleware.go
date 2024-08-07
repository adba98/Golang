package middleware

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Start")
	result := operation(add)(2, 3)
	fmt.Println(result)

	result = operation(subtract)(2, 3)
	fmt.Println(result)

	result = operation(multiply)(2, 3)
	fmt.Println(result)
}
func operation(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Star operation")
		return f(a, b)
	}
}
