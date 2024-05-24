package main

import "fmt"

// structs - are typed collection of fields
// used for grouping data together to form records

// creates person struct type that has name and age fields
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name
func newPerson(name string) *person {

	// you can safely return a pointer to local variable as a local variable will survive the scope of the function
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// this creates a new struct
	fmt.Println(person{"Bob", 20})

	// you can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 24})

	// omitted fields will be zero-values
	fmt.Println(person{name: "Kendra"})

	// an &prefix yields a pointer to the struct
	fmt.Println(&person{name: "Kimmi", age: 22})

	fmt.Println(newPerson("Alison"))

	s := person{name: "Amy", age: 35}
	fmt.Println("Name: ", s.name, "\tage: ", s.age)
}
