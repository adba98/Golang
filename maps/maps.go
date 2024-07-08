package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)
	// fmt.Println(countries)

	countries["Colombia"] = "Bogota"
	countries["Argentina"] = "Buenos aires"
	countries["Brasil"] = "Brasilia"
	/*
		fmt.Println(countries)
		fmt.Println(countries["Brasil"])
	*/

	championship := map[string]int{"Barcelona": 29, "Real Madrid": 32, "Atletico Madrid": 24}
	fmt.Println(championship)

	for team, score := range championship {
		fmt.Printf("%s team, has a score of: %d \n", team, score)
	}

	delete(championship, "Real Madrid")
	fmt.Println(championship)

	score, exist := championship["Juventus"]
	fmt.Println(score, exist)
}
