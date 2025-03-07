package main

import (
	"fmt"
)

func main() {
	// structs.go
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	fmt.Println(&Person{name: "John", age: 40})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Fred"})

	n1 := newPerson("Rohit", 37, Vertex{55, 15})
	n1.ChnageAge(29)
	fmt.Println(n1)
	n2 := NNewPerson("George")
	fmt.Println(n2)

	fmt.Println(n2.Vertex.X)
	fmt.Println(n2.Y)
	// structs fields
	v := Vertex{1, 2}
	fmt.Println("X before:", v.X)
	v.X = 4
	fmt.Println("X after:", v.X)
	// need atleast one float value on RHS so that output is 2.5, otherwise 5/2 = 2
	var y float64 = 5 / 2.00
	fmt.Println(y)
}
