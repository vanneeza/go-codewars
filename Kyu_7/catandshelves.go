package Kyu_7

func Cats(start, finish int) int {
	var jump = finish - start
	return jump/3 + jump%3
}
