package Kyu_8

import "fmt"

func Invert(arr []int) []int {
	result := make([]int, len(arr))
	for i, n := range arr {
		result[i] = -n
	}
	fmt.Println(result)
	return result
}
