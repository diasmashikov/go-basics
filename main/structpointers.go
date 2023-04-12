package main

import (
	"fmt"
) 

type Vertex struct {
	X, Y float64
}


//////////////////////////////
	

func pointervsnopointerpass() {
	v := Vertex{3, 4}
	fmt.Printf("IN MAIN %p\n", &v)
	ScaleWithPointer(&v, 10)
	fmt.Println(Sum(v))
}

func Sum(v Vertex) float64 {
	return v.X+v.Y
}


// If we pass Object not as a pointer, then it will 
// copy a whole new object without referencing the initial one
func ScaleWithPointer(v *Vertex, f float64) {
	fmt.Printf("IN METHOD %p\n", &v)
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodExtensionAndPointerReceivers() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Sum())
}

// If simple receiver, then you extend this method
func (v Vertex) Sum() float64 {
	return v.X + v.Y
}

// If *Vertex(Pointer receiver), then values of an object changes by reference and method extended to the struct
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Sum())
}