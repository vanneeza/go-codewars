package main

import (
	"github.com/vanneeza/go-codewars/Kyu_6"
	"github.com/vanneeza/go-codewars/Kyu_7"
	"github.com/vanneeza/go-codewars/Kyu_8"
)

// ///////---------------------------------
//

func main() {
	// Kyu 8 Beginner Series
	Kyu_8.Past(0, 1, 1)

	// Kyu 8 Invert
	q := []int{-5, 2, -3}
	Kyu_8.Invert(q)

	// Kyu 8 EvenOrOdd
	Kyu_8.EvenOrOdd(10)

	// Kyu 8 Reversed Strings
	Kyu_8.Solution("world")

	//----------------------------------------------------------------------------------
	// Kyu_7 : Make A function
	Kyu_7.Arithmetic(5, 7, "add")

	//------------------------------------------------------------------------------------

	// Kyu 6 : Take A ten
	roads := []rune{'n', 's', 'e', 'w', 's', 'w', 'e', 'n', 'w', 's'}
	Kyu_6.IsValidWalk(roads)

	// Kyu 6 : Help The Bookstore
	ListArt := []string{"BBART 0", "BKWRK 0", "RTSQZ 0", "DDXEF 0", "XRTYM 0"}
	ListCat := []string{"B", "R", "D", "X"}
	Kyu_6.StockList(ListArt, ListCat)

}
