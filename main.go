package main

import (
	"fmt"
	"pacote/1.tipos"
	"pacote/2.var_const"
	"pacote/3.zero_values"
	"pacote/4.funcoes"
	"pacote/5.listas"
	"pacote/6.maps"
	"pacote/7.structs"
	"pacote/8.if_else"
	"pacote/9.switch_case"
	"pacote/10.loops"
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
	listas.Show()

	fmt.Println(strings.Split("Maps", ""))
	maps.Show()

	fmt.Println(strings.Split("Structs", ""))
	structs.Show()

	fmt.Println(strings.Split("If Else", ""))
	ifelse.Show()

	fmt.Println(strings.Split("Switch Case", ""))
	switchcase.Show()

	fmt.Println(strings.Split("Loops", ""))
	loops.Show()
}