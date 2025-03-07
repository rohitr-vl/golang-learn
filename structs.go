package main

type Vertex struct {
	X int
	Y int
}
type Person struct {
	name string
	age  int
	Vertex
}

func NNewPerson(name string) *Person {
	age := 23
	v := Vertex{5, 15}
	return newPerson(name, age, v)
}

// accepts JSON obj in form of string & returns pointer to new object
func newPerson(nameStr string, ageInt int, v Vertex) *Person {
	return &Person{
		name:   nameStr,
		age:    ageInt,
		Vertex: v,
	}

}

func (p *Person) ChnageAge(age int) {
	p.age = age
}
