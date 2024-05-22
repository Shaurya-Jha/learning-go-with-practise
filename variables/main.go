package main

import (
	"fmt"
	"reflect"
)

func main() {
	// single variable declaration
	var a = "initial"
	fmt.Println(a)

	// multiple variables declaration
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go automatically infers the variable type
	var d = true
	fmt.Println(d, "Type of variable d is: ", reflect.TypeOf(d))

	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing a variable
	f := "apple"
	fmt.Println("Value of variable f is: ", f)
}
