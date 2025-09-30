package structs

import "fmt"

type Animal struct {
	Nome string
	Idade int
// campo nome e campo idade
}

type Pessoa struct {
	Animal
	Tipo string
}

func Show() {
	fmt.Println(Animal{"Cachorro", 5})
	fmt.Println(Animal{Nome: "Cachorro", Idade: 5})
	fmt.Println(Animal{Nome: "Cachorro"}) // Idade não precisa ser informada, porém terá o valor de zero value

	dog := Animal{Nome: "Cachorro", Idade: 5}
	fmt.Println(dog.Nome)
	fmt.Println(dog.Idade)

	dog.Idade = 10
	fmt.Println(dog.Idade)

	cat := Animal{"Gato", 3}
	fmt.Println(cat)

	animais := []Animal{}
	animais = append(animais, dog, cat)
	fmt.Println(animais)

	// structs + maps
	selva := map[string][]Animal{}
	selva["Safari"] = animais
	fmt.Println(selva)

	domesticos := map[string][]Animal{}
	domesticos["Casa"] = animais
	fmt.Println(domesticos)

	var animal2 = map[string][]Animal{
		"Petshop": {{Nome: "Papagaio", Idade: 2}, {Nome: "Tartaruga", Idade: 80}},
		"Fazenda": {{Nome: "Vaca", Idade: 5}, {Nome: "Cavalo", Idade: 7}},
	}

	fmt.Println(animal2)

	pessoa := Pessoa{dog, "Humano"}
	fmt.Println(pessoa)
	fmt.Println(pessoa.Animal.Nome)
	fmt.Println(pessoa.Animal.Idade)
	fmt.Println(pessoa.Tipo)

	fmt.Println()
}



/*
- criar sua própria estrutura de dados
- personalizar conforme necessidade
*/