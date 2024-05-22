package main

import "fmt"

func main() {
	// create an array that holds 5 values
	var a [5]int
	fmt.Println("Emp: ", a)

	// set the value of array at position 4
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Dcl: ", b)

	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Dcl: ", c)

	// two dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
}
