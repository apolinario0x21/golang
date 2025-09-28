package tipos

import "fmt"

func Show() {
	fmt.Printf("Tipo:  %T - Valor: %v. %v", true, true, "\n")
	fmt.Printf("Tipo: %T - Valor: %v", "Tester", "String\n")
	fmt.Printf("Tipo: %T - Valor: %v %v", 1, 1, "\n")
	fmt.Printf("Tipo: %T - Valor: %v", "2", "2\n")
	fmt.Printf("Tipo: %T - Valor: %v %v", 7.543, 9.876,"\n\n")
	
}

// String - sequência de bytes
// %v - valor padrão
// %T - tipo do valor