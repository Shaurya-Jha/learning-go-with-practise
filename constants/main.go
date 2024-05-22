package main

import (
	"fmt"
	"math"
)

// const keyword is used to declare the constant variable
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000
	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
