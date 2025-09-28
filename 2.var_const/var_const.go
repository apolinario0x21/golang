package varconst

import "fmt"

func Show() {
	var name string
	name = "tester"
	fmt.Println(name)

	name = "admin"
	fmt.Println(name)

	var age int
	age = 21
	fmt.Println(age)

	var isAdmin = "yes"
	fmt.Println(isAdmin)

	var isUser, isTester = 3, 2
	fmt.Println(isUser, isTester)

	var admin = true
	fmt.Println(admin)

	city := "New York"
	fmt.Println(city)

	const pi = 3.1415
	fmt.Print(pi, "\n\n")
}

// Inferência de tipo, o Go consegue identificar o tipo da variável