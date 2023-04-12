package main

import (
	"math"
)

type Vertex struct {
	X, Y float64
}

// Struct extension or receivers
// You cann access variables of a struct 
// through the receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// }