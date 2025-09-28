package main

import (
	"fmt"
	"pacote/1.tipos"
	"strings"
	"pacote/2.var_const"
	"pacote/3.zero_values"
)

func main() {
	fmt.Print("Pacotes fmt e strings, função Println e Split\n\n")
	fmt.Println(strings.Split("Tipos", ""))
	tipos.Show()	
	fmt.Println(strings.Split("Var and Const", ""))
	varconst.Show()
	fmt.Println(strings.Split("Zero Values", ""))
	zero_values.Show()
}