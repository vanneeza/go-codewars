package Kyu_8

/**
Create a function that takes an integer as an argument and returns "Even"
for even numbers or "Odd" for odd numbers.

*/

func EvenOrOdd(number int) string {
	result := ""
	if number%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	return result
}
