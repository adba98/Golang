package variables

import "fmt"

func ShowIntegers() {
	var intCommon int
	intFrom32 := int32(10)
	intFrom64 := int64(10)

	fmt.Println("int common:", intCommon)
	fmt.Println("int 32:", intFrom32)
	fmt.Println("int 64:", intFrom64)
}
