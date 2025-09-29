package main

import (
	"fmt"
	"pacote/1.tipos"
	"pacote/2.var_const"
	"pacote/3.zero_values"
	"pacote/4.funcoes"
	"pacote/5.listas"
	"strings"
)

func main() {
	fmt.Print("Pacotes fmt e strings, função Println e Split\n\n")
	fmt.Println(strings.Split("Tipos", ""))
	tipos.Show()	
	fmt.Println(strings.Split("Var and Const", ""))
	varconst.Show()
	fmt.Println(strings.Split("Zero Values", ""))
	zero_values.Show()
	fmt.Println(strings.Split("Funções", ""))
	funcoes.Show()

	fmt.Println(strings.Split("Listas", ""))
	var arr [2]string
	arr[0] = "Hi"
	fmt.Println(arr)
	
	listas.Show()
	



}