package main

import (
	"fmt"
	"godesde0/interfaces"
	"godesde0/models"
)

func main() {
	Pedro := new(models.Man)
	HumanBreathes(Pedro)

	Ana := new(models.Women)
	HumanBreathes(Ana)
}

func HumanBreathes(hu interfaces.Human) {
	hu.Breathe()
	fmt.Printf("I am a %s, and i am breathes \n", hu.Sex())
}
