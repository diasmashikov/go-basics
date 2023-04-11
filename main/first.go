package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var president = "Tokaev"

func main() {
	
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