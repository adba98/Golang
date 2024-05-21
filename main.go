package main

import (
	"fmt"
	"godesde0/variables"
)

func main() {
	state, text := variables.CovertToText(111)
	fmt.Println(state)
	fmt.Println(text)
}
