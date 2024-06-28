package arrays

import "fmt"

var table2 []int = []int{2, 5, 4}
var array [10]int = [10]int{6, 3, 7, 44, 58, 73}

func ShowSlices() {
	fmt.Println(table2)

	slice := array[3:]
	fmt.Println(slice)

	slice2 := array[:5]
	fmt.Println(slice2)

	slice3 := array[3:6]
	fmt.Println(slice3)
}

func Capacity() {
	elements := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacity %d", len(elements), cap(elements))

	numbers := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i)
	}
	fmt.Printf("\nLargo %d, Capacity %d", len(numbers), cap(numbers))
}
