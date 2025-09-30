package ifelse

import "fmt"

func Show() {
	valor := 10

	if valor == 7 {
		fmt.Println("Valor é 1")
	} else if valor == 2 {
		fmt.Println("Valor não é 2")
	} else {
		fmt.Println("Valor não é 1 ou 2, é outro")
	}

	fmt.Println()
}