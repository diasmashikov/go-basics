package main

import (
	"fmt"
	"math/rand"
	"time"
)



func main() {
	
}

func makeSequence() func() int {
	i:=1
	return func() int {
	   i+=2
	   return i
	}
 }
 

func sum(numbers ...int) {
	total := 0
    for _, num := range numbers {
        total += num
    }
    fmt.Println(total)
}

func deferMethod() {
	defer who()
	hello()
}

func hello() {
	fmt.Println("Hello")
}

func who() {
	fmt.Println("WHO")
}


func pointers() {
	var x int = 5

	var ipointer *int

	ipointer = &x 

	fmt.Printf("Address of variable x: %x\n", &x)
	fmt.Printf("Storage for memory address is pointer %x\n", ipointer)
	fmt.Printf("Value of variable x: %d\n", *ipointer)
}



func diceRoll() {
	rand.Seed(time.Now().UnixNano())
	n := 10
	var dice int = 6
	for i := 0; i < n; i++ {
		fmt.Printf("ROll THE DICE %d\n", rand.Intn(dice) + 1)
	}
}

func randomNumbers() {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(10)
	randomFloat := rand.Float64()
	fmt.Println(randomInt, randomFloat)
}

func Mappa() {
	elements := make(map[string]string)
    elements["O"] = "Oxygen"
    elements["Ca"] = "Calcium"
    elements["C"] = "Carbon"

    fmt.Println(elements["C"])

	alpha := map[string]int{
		"A" : 5,
		"B" : 4,
		"C" : 3,
	}
	fmt.Println(alpha["A"])
}

func structInGo() {
	var aperson Person 

	aperson.name = "Albert"
	aperson.job = "Professor"

	fmt.Printf("aperson.name = %s\n", aperson.name)
	fmt.Printf("aperson.job = %s\n", aperson.job)

	house := House{
		"Mangilik el",
		32767,
		127,
	}

	fmt.Print(house)
}

type Person struct {
	name string 
	job string
}

type House struct {
	street string 
	price int16
	rooms int8
}