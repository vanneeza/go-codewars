/**
You live in the city of Cartesia where all roads are laid out in a perfect grid.
You arrived ten minutes too early to an appointment, so you decided to take the opportunity
to go for a short walk. The city provides its citizens with a Walk Generating App on their
phones -- everytime you press the button it sends you an array of one-letter strings
representing directions to walk (eg. ['n', 's', 'w', 'e']). You always walk only a single
block for each letter (direction) and you know it takes you one minute to traverse one city
block, so create a function that will return true if the walk the app gives you will take
you exactly ten minutes (you don't want to be early or late!) and will, of course, return
you to your starting point. Return false otherwise.

The point is
1. One block, 1 minute
2. If 10 minute give return true, so it means 10 block
likes {'n','s','e''n','s','w''n','s','w','e'}
3. Return false!
- return to your starting point is false
- more than 10 block give return false
- less than 10 block give return false

What should we do ?

1. using function len() to get a lenght from variable walk
if lenght not 10, I will give return false
2. Create variable x,y int, we set 0,0
3. using for to get lenght from walk
4, inside for we create switch case and dont forget to convert rune to string

*/

package Kyu_6

import "fmt"

func IsValidWalk(walk []rune) bool {
	// len(walk) for check lenght from variable walk
	if len(walk) != 10 {
		return false
	}

	x, y := 0, 0
	for _, z := range walk {
		walking := string(z) // string() to convert rune to string
		switch walking {
		case "n":
			y++
		case "s":
			y--
		case "e":
			x++
		case "w":
			x--
		}
		fmt.Println(walking)
	}
	return x == 0 && y == 0
}
