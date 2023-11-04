package main

import "fmt"

func main() {
	// variable
	// don't 1One, special character, avoid build in function name
	// do string only, exception num1
	// standardization
	// don't start char capitalize, num_one (snake case), numone
	// do numOne camel case

	// var  name data type
	var numOne int
	fmt.Println(numOne)

	numOne = -129
	fmt.Println(numOne)

	var numTwo int32 = 1333
	fmt.Println(numTwo)

	var numThree float32
	fmt.Println(numThree)

	numFour := 2.4
	fmt.Println(numFour)

	var one, two, three int
	one, two, three = 1, 2, 3
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)

	var four, five, six int = 4, 5, 6
	fmt.Println(four)
	fmt.Println(five)
	fmt.Println(six)

	seven, eight, nine := 7, 8, 9
	fmt.Println(seven)
	fmt.Println(eight)
	fmt.Println(nine)

	ten, hello, condition := 10, "hello world", true
	fmt.Println(ten, hello, condition)

	world := "world"
	var newStr string = world + " " + hello
	fmt.Println(newStr)

	sum := numFour + ten
	fmt.Println(sum)
}
