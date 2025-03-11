package main

import (
	"fmt"
	"math"
)

type rect1 struct {
	width, height float64
}

func (r rect1) area() float64 {
	return r.width * r.height
}

func (rp *rect1) perimeter() float64 {
	return 2*rp.width + 2*rp.height
}

func getRectAreaPerimemter() {
	rectObj := rect1{3, 4}
	rectObjPointer := &rectObj

	// both the methods can be called with rectObj & rectObjPointer
	fmt.Println("Area: ", rectObj.area())
	fmt.Println("Perimeter: ", rectObjPointer.perimeter())
}

type VertexFloat struct {
	X, Y float64
}

// pointer receivers in Methods
func (v VertexFloat) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VertexFloat) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
