package main

import (
	"github.com/vanneeza/go-codewars/Kyu_6"
	"github.com/vanneeza/go-codewars/Kyu_7"
	"github.com/vanneeza/go-codewars/Kyu_8"
)

func main() {
	// ---------------------------------------------K Y U U : 8 ------------------------------
	//	Beginner Series
	Kyu_8.Past(0, 1, 1)

	// Invert
	q := []int{-5, 2, -3}
	Kyu_8.Invert(q)

	// EvenOrOdd
	Kyu_8.EvenOrOdd(10)

	// Reversed Strings
	Kyu_8.Solution("world")

	// //--------------------------K Y U U :  7--------------------------------------------------
	Kyu_7.Arithmetic(5, 7, "add")
	Kyu_7.Cats(1, 5)

	// //--------------------------K Y U U :  6---------------------------------------------------

	//  Take A ten
	roads := []rune{'n', 's', 'e', 'w', 's', 'w', 'e', 'n', 'w', 's'}
	Kyu_6.IsValidWalk(roads)

	// Help The Bookstore
	ListArt := []string{"BBART 0", "BKWRK 0", "RTSQZ 0", "DDXEF 0", "XRTYM 0"}
	ListCat := []string{"B", "R", "D", "X"}
	Kyu_6.StockList(ListArt, ListCat)

}
