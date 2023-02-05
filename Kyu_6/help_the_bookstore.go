/**
A bookseller has lots of books classified in 26 categories labeled A, B, ... Z.
Each book has a code c of 3, 4, 5 or more characters. The 1st character of a code is a
capital letter which defines the book category.

In the bookseller's stocklist each code c is followed by a space and by a positive
integer n (int n >= 0) which indicates the quantity of books of this code in stock.
For example an extract of a stocklist could be:

L = {"ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"}.
or
L = ["ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"] or ....

You will be given a stocklist (e.g. : L) and a list of categories in capital letters e.g :

M = {"A", "B", "C", "W"}
or
M = ["A", "B", "C", "W"] or ...

and your task is to find all the books of L with codes belonging to each
category of M and to sum their quantity according to each category.
For the lists L and M of example you have to return the string
(in Haskell/Clojure/Racket/Prolog a list of pairs):

(A : 20) - (B : 114) - (C : 50) - (W : 0)

where A, B, C, W are the categories, 20 is the sum of the unique book
of category A, 114 the sum corresponding to "BKWRK" and "BTSQZ",
50 corresponding to "CDXEF" and 0 to category 'W' since there are no code beginning with W.
If L or M are empty return string is "" (Clojure/Racket/Prolog should return an empty
array/list instead).

Notes:
In the result codes and their values are in the same order as in M.
See "Samples Tests" for the return.

The Point Is ?

- find all the books of L with codes belonging to each
category of M and to sum their quantity according to each category.

- The result must be like (A : 20) - (B : 114) - (C : 50) - (W : 0)

- If L or M are empty return string is ""

So, What should we do ?

1. I want to split array L = {"ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"} to be
label and qty and take take first character.

(e.g) L = {A 20, C 50 , B 25, B 89, D 60}

2. Add array M {A B C W} to L, if the value of 2 array have the same characters don't add. Just
add the different characters. Of course, if the array have same characters sum the qty
In this case, like L = {A 20, B 144, C 50, W 0}

3. Return the value! :v

*/

package Kyu_6

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {

	var listBook []string     // Making variable slice listbook
	var result []string       // make variable slice result
	m := make(map[string]int) // make variable maps m

	for _, v := range listArt { // looping value listArt and take the value

		label := v[0:1] // Create Label, v[0:1] for take the first letter from v

		// Create Qty, using regexp.MustCompile for find the number from string
		// .FindString will search for the first string that matches the regular expression in the variable v.

		qty := regexp.MustCompile("[0-9]+").FindString(v)
		listBook = append(listBook, label+" "+qty)
		//  append label and qty to slice string listbook.
		// with space (" ") between the value
	}

	for _, v := range listBook { // Looping slice listbook for take the value
		s := strings.Split(v, " ")   // split string v using space  as the separator
		label := s[0]                // take the first character
		qty, _ := strconv.Atoi(s[1]) // convert to int
		m[label] += qty              // increments the value stored at the key label in the m map by the converted
		//integer quantity. The m map is used to store the total quantity of books for each label.
	}

	// if !ok statement then checks if the value of ok is false, which means
	// that m[v] does not exist in the map, and the key-value pair v:0 is added to the map.
	for _, v := range listCat {
		if _, ok := m[v]; !ok {
			m[v] = 0
		}
	}

	for _, v := range listCat {
		result = append(result, "("+v+" : "+strconv.Itoa(m[v])+")") // append to slice result
	}

	resultBanget := strings.Join(result, " - ") // string.join for create new string
	// using " - " as separator

	fmt.Println(resultBanget)
	if len(listArt) == 0 || len(listCat) == 0 { // if lenght listart and list cat 0
		return "" // return ""
	}

	return resultBanget
}

/**
This is my code to pass the test.
On the solution page. I see the best practice to pass this test
if you want to see. I will give.

package kata

import (
  "strings"
  "strconv"
  "fmt"
)

func StockList(listArt []string, listCat []string) string {
    if len(listArt) == 0 || len(listCat) == 0 {
      return ""
    }

    quantity := make(map[string]int)
    for _, stock := range listArt {
      splitted := strings.Split(stock, " ")
      firstCh := string(splitted[0][0])
      n, _ := strconv.Atoi(splitted[1])
      quantity[firstCh] += n
    }

    result := make([]string, len(listCat))
    for i, code := range listCat {
      result[i] = fmt.Sprintf("(%s : %v)", code, quantity[code])
    }

    return strings.Join(result, " - ")
}

Very simple hahhahahaa, thxu
*/
