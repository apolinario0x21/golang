package funcoes

import "fmt"

func soma(x, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func name(nome string) string {
	return nome
}

func Show() {
	sum := soma(7, 3)
	fmt.Println(sum)

	sub := sub(7, 3)
	fmt.Println(sub)

	nome := name("Tester")
	fmt.Println(nome)
}


// funções com letra minúsca são privadas ao pacote
// funções com letra maiúscula são públicas ao pacote