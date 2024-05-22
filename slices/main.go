package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("Un-initialized: ", s, s == nil, len(s) == 0)

	// create an empty slice with non-zero length i.e 3
	s = make([]string, 3)
	fmt.Println("Empty: ", s, "length: ", len(s), "cap: ", cap(s))

	// set the values in slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	// get values from slice s
	fmt.Println("get: ", s[2])

	// return the length of slice
	fmt.Println("length : ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Apd: ", s)

	// make another slice named c and copy s slice to c
	c := make([]string, len(s))
	// copy from s to c slice
	copy(s, c)
	fmt.Println("Copy: ", c)

	l := s[2:5]
	fmt.Println("Slice 1: ", l, "Length: ", len(l))

	l = s[:5]
	fmt.Println("Slice 2: ", l, "Length: ", len(l))

	l = s[2:]
	fmt.Println("Slice 3: ", l, "Length: ", len(l))

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
}
