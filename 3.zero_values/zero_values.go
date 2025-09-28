package zero_values

import "fmt"

func Show() {

	var i int
	var f float64
	var f2 float32
	var b bool
	var s string
	
	fmt.Printf("%T - %v \n%T - %v\n%T - %v \n%T - %v \n%T - %q \n ", i, i, f, f, f2, f2, b, b, s, s)
}

// %q - string com aspas