package listas

import ( "fmt" )

var a [2]string
	a[0] = "Hello"
	a[1] = "World"

var array = [2]string{"Hello", "CodeSpaces"}


func Show() {
	fmt.Println(array)
	fmt.Println(array[0], array[1])
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

Maps: Heterogêneos
Chave e valor podem ser de tipos diferentes
{"nome": "Hulk", "idade": 300} = map[string]



*/