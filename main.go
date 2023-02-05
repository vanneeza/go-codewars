package main

import (
	Kyu_6 "github.com/vanneeza/go-codewars/Kyu_6"

	makeafunction "github.com/vanneeza/go-codewars/Kyu_7"
)

// ///////---------------------------------

func main() {

	// Make A function
	makeafunction.Arithmetic(5, 7, "add")

	// Kyu 6 : Take A ten
	roads := []rune{'n', 's', 'e', 'w', 's', 'w', 'e', 'n', 'w', 's'}
	Kyu_6.IsValidWalk(roads)

	// Kyu 6 : Help The Bookstore
	ListArt := []string{"BBART 0", "BKWRK 0", "RTSQZ 0", "DDXEF 0", "XRTYM 0"}
	ListCat := []string{"B", "R", "D", "X"}
	Kyu_6.StockList(ListArt, ListCat)

}
