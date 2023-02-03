package makeafunction

import (
	"fmt"
	"math"
)

func Arithmetic(a int, b int, operator string) int {
	ChangePositiveA := int(math.Abs(float64(a)))
	ChangePositiveB := int(math.Abs(float64(b)))

	var r float64
	switch operator {
	case "add":
		r = float64(ChangePositiveA) + float64(ChangePositiveB)
	case "subtract":
		r = float64(ChangePositiveA) - float64(ChangePositiveB)
	case "divide":
		r = float64(ChangePositiveA) / float64(ChangePositiveB)
	case "multiply":
		r = float64(ChangePositiveA) * float64(ChangePositiveB)
	default:
		fmt.Println("Error: invalid operator:", operator)
	}
	fmt.Println("Result:", r)
	ConvertToInt := int(r)
	return ConvertToInt
}
