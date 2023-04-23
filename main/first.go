package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var president = "TokaevPresidentKzAstana"


func interacesInGo() {

}

func writeToFile() {
	file, err := os.Create("main/random-file.txt")

    if err != nil {
        return
    }
    defer file.Close()

	data := []byte("Sharrrap guyzzz!!!!")

    file.Write(data)

	var idx int64 = int64(len(data))

	data2 := []byte("Sharrrap guyzzz!!!!")

	file.WriteAt(data2, idx)
}

func readFileWithLines() {
	lines, err := readLines("main/random-file.txt")
  	if err != nil {
    	log.Fatalf("readLines: %s", err)
  	}
  	// print file contents
  	for i, line := range lines {
    	fmt.Println(i, line)
  	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
	  return nil, err
	}
	defer file.Close()
  
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	  lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
  }

func readWholeFile() {
	b, err := ioutil.ReadFile("main/random-file.txt")

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	fmt.Println(str)
}

func fileExists() {
	if _, err := os.Stat("main/random-file.txt"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n");
	}
}

func switchStatement() {
	switch a := 1; {
    case a == 2:
        fmt.Println("The integer was == 2")
    case a == 5:
        fmt.Println("The integer was == 5")
		fallthrough
    case a == 3:
        fmt.Println("The integer was == 3")
        fallthrough
    case a == 4:
        fmt.Println("The integer was == 4")
    case a == 1:
        fmt.Println("The integer was == 1")
    default:
        fmt.Println("default case")
    }
}

func whileLoop() {
	i := 1 
	max := 20 

	for i < max {
		fmt.Println(i)
		i++
	}
}

func ifStatements() {
	var age = 16

	if (age >= 18) {
		fmt.Println("You can drink")
	} else if age >= 16 && age < 18 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You are a kid bruvvv")
	}
}

func rangeLoop() {
	nums := []int{10, 20, 30, 40, 50, 60}
	
	for i, num := range nums {
		fmt.Printf("Index: %d, Number: %d\n", i, num)
	}
}

func basicForLoops() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	slice := []int{10, 20, 30, 40, 50}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Lamadre %d", i)
		fmt.Println()
		for j := 0; j < 10; j++ {
			fmt.Printf("Bambino %d ", j)
		}
		fmt.Println()
	}
}


func slicesInGoThree() {
	a := make([]int, 5)
	printSliceMake("a", a)

	b := make([]int, 0, 5)
	printSliceMake("b", b)

	c := b[:2]
	printSliceMake("c", c)

	d := c[2:5]
	printSliceMake("d", d)
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func slicesInGoTwo() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


func slicesInGoOne() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[2:4]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func arraysInGo() {
	var arr1 = new([5]int)
	fmt.Println(arr1[0])
	var arr2 = [5]int{18, 20, 15, 22, 16}
	fmt.Println(arr2[0])
	var arrKeyVal = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrKeyVal[3])
}

func scopingDemonstration() {
	fmt.Println(president)
	var president = "Dias"
	fmt.Println(president)
}

func multipleVariablesAndSwap() {
	a, b, c := 5, 7, "abc"
	fmt.Println(a, b, c)
	a, b = b, a
	fmt.Println(a, b, c)
	fmt.Println(president)
}

func discountedPrice() {
	applePrice := 3
	appleDiscount := 0.1
	finalApplePrice := float64(applePrice) - float64(applePrice) * appleDiscount
	fmt.Printf("The final price is %.1f", finalApplePrice)
}

func heightAndAge() {
	reader := bufio.NewReader(os.Stdin)
	
	age := enterDigitInfo("age", reader)
	height := enterDigitInfo("height", reader)

	fmt.Printf("You are %d years old and height of %d cm", age, height)
}

func enterDigitInfo(attribute string, reader *bufio.Reader) int {
	fmt.Printf("Enter your %s: ", attribute)
	return inputDigit(reader)
}

func inputDigit(reader *bufio.Reader) int {
	strDigit, _ := reader.ReadString('\n')
	digit, _ := strconv.Atoi(strings.TrimSpace(strDigit))
	return digit
}

func fullname() {
	var firstName string = "Dias"
	var lastName string = "Mashikov"
	fullname := strings.Join([]string{firstName, lastName}, " ")
	fmt.Printf("%s", fullname)
}