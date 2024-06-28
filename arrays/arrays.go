package arrays

import "fmt"

var table [10]int = [10]int{10, 0, 59}
var matrix [5][10]int

func ShowArrays() {
	table[7] = 33
	table[2] = 54
	fmt.Println(table)

	table2 := [10]string{"Pedro", "Juan"}
	fmt.Println(table2)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}

	matrix[2][8] = 15
	fmt.Println(matrix)

}
