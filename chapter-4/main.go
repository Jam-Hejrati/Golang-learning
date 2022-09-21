package main

import "fmt"

// Variables
var globallyVariable = "this is a global variable" // type 'string' could omit so omitted :)

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	var y string
	y = "Hello Jam"
	fmt.Println(y)
	y = "Hello Jam 2"
	fmt.Println(y)
	y = y + " Second"
	fmt.Println(y)
	z := " this is a string"
	w := 45
	fmt.Println(z, w)
	fmt.Println(globallyVariable)

	const str string = "first" // can't change the value anymore
	fmt.Println(str)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a + b + c)
}
