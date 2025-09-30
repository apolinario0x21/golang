package loops

import "fmt"

func Show() {
	sum := 0
	for i := 0; i <= 4; i++ {
		sum += i
	}

	fmt.Println("Soma:", sum)

	carros := []string{"Fusca", "Brasília", "Chevette"}

	for _, carro := range carros {
		fmt.Println(carro)
	}
}