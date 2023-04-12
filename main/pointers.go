package main

import "fmt"

type Point struct {
	X, Y int
}

var (
	v1 = Point{1, 2}  // has type Vertex
	v2 = Point{X: 1}  // Y:0 is implicit
	v3 = Point{}      // X:0 and Y:0
	p  = &Point{1, 2} // has type *Vertex
)

func structLiterals() {
	fmt.Println(v1, p, v2, v3)
	fmt.Println(*p)
}

func structPointers() {
	v := Point{1, 2}
	p := &v
	// Automatic deference P.X = value
	// instead of 
	// *p.X = 1000
	p.X = 1000
	fmt.Println(v)
}


func basicsOfPointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Printf("Pointer of i: %x\n", p)
	fmt.Printf("Reading i by denoting pointer p: %d\n", *p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Printf("Modifying i value by denoting pointer p and modifying: %d\n", i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Printf("Modified j without mentioning it by working with p pointer and denotation %d", j) // see the new value of j
}