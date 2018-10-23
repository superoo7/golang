package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Integer
	var a int8 // uint8
	var b int16 // uint 16
	var c int32 // uint 32
	var d int64 // uint 64
	fmt.Println("Init value for int is ", a)
	a = 127
	b = 15
	c = 15
	d = 15
	fmt.Println("Value of a is ", a)
	fmt.Println("Value of b is ", b)
	fmt.Println("Value of c is ", c)
	fmt.Println("Value of d is ", d)
	// Boolean
	var booleanTrue bool
	fmt.Println("Init value for boolean is ", booleanTrue)
	booleanTrue = true
	fmt.Println(booleanTrue)
	// String
	var stringName string
	fmt.Println("Init value for string is ", stringName)
	stringName = "string stuff"
	fmt.Println(stringName)
	// Float
	var floatNum float32 // float64
	fmt.Println("Init value for float is ", floatNum)
	floatNum = 1.12
	fmt.Println(floatNum)
	// Complex
	var complexNum complex64 // complex128S
	fmt.Println("Init value for complex is ", complexNum)
	complexNum = (1 + 3i)/(3i)
	complexAng := cmplx.Acos(-2 + 12i)
	fmt.Println(complexNum)
	fmt.Println(complexAng)

	}
