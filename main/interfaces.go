package main

import (
	"fmt"
	"math"
)

// If a method name is the same as the interface implementation method
// THEN IT AUTOMATICALLY EXTENDS THAT INTERFACE!
type Abser interface {
	Abs() float64
}


func interfaceDemo() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Dot{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser
	fmt.Println(a)

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Dot struct {
	X, Y float64
}

func (v *Dot) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


////

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}