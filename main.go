package main

import (
	"fmt"
	"godesde0/goroutines"
)

func main() {
	go goroutines.MySlowName("Oscar")
	fmt.Println("Enter number:")
	var x string
	fmt.Scanln(&x)

}
