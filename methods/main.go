package main

import "fmt"

// rect struct
type rect struct {
	width, height int
}

// methods can be defined for either pointer receiver type or value receiver types

// area method with pointer receiver
func (r *rect) area() int {
	return r.width * r.height
}

// perimeter method with value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Ares: ", r.area())
	fmt.Println("Perimeter: ", r.perim())
}
