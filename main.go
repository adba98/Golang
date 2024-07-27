package main

import (
	"fmt"
	"godesde0/goroutines"
)

func main() {
	chanel := make(chan bool)
	go goroutines.MySlowName("Oscar", chanel)
	<-chanel
	fmt.Println("Here")
}
