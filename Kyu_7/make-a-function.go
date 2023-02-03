/**
Given two numbers and an arithmetic operator (the name of it, as a string),
return the result of the two numbers having that operator used on them.
a and b will both be positive integers, and a will always be the first number
in the operation, and b always the second.

The four operators are "add", "subtract", "divide", "multiply".

A few examples:(Input1, Input2, Input3 --> Output)

5, 2, "add"      --> 7
5, 2, "subtract" --> 3
5, 2, "multiply" --> 10
5, 2, "divide"   --> 2.5

Try to do it without using if statements!

What will i do ?
1. Make 3 variable, A & B int and Operator string
2. "a and b will both be positive integers", I think i will give
	a functions for var A & B, so they are always positive and can't minus
	using int(math.Abs()). You must import package "math" before using this functions
3. "A will always be the first number in the operation, and B always the second"
	don't forget about the placement like A + B,
4 "without using if statements". I will using switch case because it very simple. The Operators
	is string. So, with switch case we can do it like
	case "add" {
		a + b
	}
*/

package makeafunction

import (
	"fmt"
	"math"
)

func Arithmetic(a int, b int, operator string) int {

	ChangePositiveA := int(math.Abs(float64(a))) // for convert min numbers (-10) to (10)
	ChangePositiveB := int(math.Abs(float64(b)))

	var r float64
	switch operator {
	case "add":
		r = float64(ChangePositiveA) + float64(ChangePositiveB) // float64 for used to represent decimal values in mathematical
	case "subtract":
		r = float64(ChangePositiveA) - float64(ChangePositiveB)
	case "divide":
		r = float64(ChangePositiveA) / float64(ChangePositiveB)
	case "multiply":
		r = float64(ChangePositiveA) * float64(ChangePositiveB)
	default:
		fmt.Println("Error: invalid operator or using lowercase, hehe", operator)
	}
	fmt.Println("Result:", r) // To print line the result
	ConvertToInt := int(r)    // Convert Variable R (float64) to Int, because we want return is int
	return ConvertToInt

	/** This is my code to pass the test.
	On the solution page. I see the best practice to pass this test
	if you want to see. I will give.

	func Arithmetic(a int, b int, operator string) int{
		switch operator {
		 case "add":
			return a + b
		  case "subtract":
			return a - b
		  case "multiply":
			return a * b
		  case "divide":
			return a / b
		  default:
			panic("Unknown operator")
		}
	  }

	  This code is very simple! Thxu...
	*/
}
