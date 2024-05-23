package main

import "fmt"

// zeroval has an int parameter, so args will be passed to it by value
func zeroval(ival int) {
	ival = 0
}

// zeroptr has an *int parameter, meaning that it takes an int pointer
func zeroptr(iptr *int) {
	// the *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	// the &i syntax gives the memory address of i i.e a pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	// pointers can be printed too
	fmt.Println("pointer: ", &i)
}
