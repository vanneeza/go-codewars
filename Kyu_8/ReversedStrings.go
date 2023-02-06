package Kyu_8

import "fmt"

func Solution(word string) string {
	var result = ""
	for _, c := range word {
		result = string(c) + result
	}
	fmt.Println(result)
	return result
}
