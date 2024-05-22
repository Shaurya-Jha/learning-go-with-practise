package main

import "fmt"

func main() {
	i := 1

	fmt.Println("For loop:")
	// basic for loop with one condiion
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// classical initial/condition/after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// basic do this N times iteration
	for i := range 3 {
		fmt.Println("Range: ", i)
	}

	// for loop withour condition
	for {
		fmt.Println("loop")
		break
	}

	// continue to the next iteraion of the loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("\n\nIf else loop: ")
	// basic if and else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// you can have if statement without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Print(num, " is negative")
	} else if num > 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has numtiple digits")
	}
}
