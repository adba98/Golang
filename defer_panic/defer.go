package defer_panic

import (
	"fmt"
	"log"
)

func Defer() {
	fmt.Println("First message")
	defer fmt.Println("Last message")

	fmt.Println("Second message")
}

func Panic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalln("Error", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Found value 1")
	}
}
