package listas

import ( "fmt" )

func ArrayDeStrings(str, str1 string) [2]string{
var arr [2]string
	arr[0] = str
	arr[1] = str1

	return arr
}

var ArrayDeNumeros = [2]string{"0", "1"}


func SliceDeStrings(str, str1 string) [5]string {
	isSlice := make([]string, 5)
	isSlice[0] = str
	isSlice[1] = str1

	return [5]string(isSlice)
}

func NumerosPares() {
	isPar := []int{10, 20, 30, 40, 50}
	fmt.Println(isPar)

	numeros := append(isPar, 60, 70, 80, 90, 100)
	fmt.Println(numeros)
}

func Show() {
	fmt.Println(ArrayDeStrings("Zero", "Um"))
	
	fmt.Println(ArrayDeNumeros)
	fmt.Println(ArrayDeNumeros[0], ArrayDeNumeros[1])

	fmt.Println(SliceDeStrings("Dez", "Vinte"))
	NumerosPares()

	fmt.Println()
}

/*
Arrays e Slices: Homogêneos
Todos os elementos são do mesmo tipo
[1, 2, 3] = []int
['a', 'b', 'c'] = []string

ARRAYS: 
- tamanho fixo
- acessamos os valores com índice
- retornar o tamanho do array com len()

SLICE
- não possui tamanho fixo
- acessamos os valores com índice
- retornar o tamanho do slice com len()
- fun append() add valor ao slice
*/