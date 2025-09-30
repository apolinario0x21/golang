package switchcase
import "fmt"

func Show() {
posicao := 1
switch posicao {
case 1:
	fmt.Println("Primeiro")
case 2:
	fmt.Println("Segundo")
case 3:
	fmt.Println("Terceiro")
default:
	fmt.Println("Fora do pódio")
}

fmt.Println()
}