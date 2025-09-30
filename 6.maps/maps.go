package maps

import "fmt"

func CreateMapSemDados() {
	linguagens := map[string]int{}
	linguagens["Java"] = 1995
	linguagens["Golang"] = 2009

	fmt.Println(linguagens)
	fmt.Println(linguagens["Java"])
	fmt.Println(linguagens["Golang"])
	
}

func CreateMapComDados() {
	carros := map[string]int {
		"BYD DOLPHIN MINI": 2023,
		"BYD SEAL": 2023,
	}

	
	fmt.Println(carros)
	fmt.Println(carros["BYD SEAL"])

	carros["BYD SHARK"] = 2024
	fmt.Println(carros)
}

func Show() {
	CreateMapSemDados()
	CreateMapComDados()

	fmt.Println()
}

/*
Maps: Heterogêneos
Chave e valor podem ser de tipos diferentes
{"nome": "Hulk", "idade": 300} = map[string]
*/