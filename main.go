package main

import (
	"fmt"
	"runtime"
)

func main() {
	if os := runtime.GOOS; os == "windows" || os == "OS X." {
		fmt.Println("Is Windows")
	} else {
		fmt.Println("Is", runtime.GOOS)
	}

	switch os := runtime.GOOS; os {

	case "linux":
		fmt.Println("You are using Linux.")
	case "windows":
		fmt.Println("You are using Windows.")
	default:
		fmt.Printf("%s is your operating system", runtime.GOOS)
	}

}
