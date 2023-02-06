package Kyu_8

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Past(h, m, s int) int {
	a := strconv.Itoa(h)
	b := strconv.Itoa(m)
	c := strconv.Itoa(s)
	d := []string{a, "h", b, "m", c, "s"}
	f := strings.Join(d, "")
	z, _ := time.ParseDuration(f)
	result, _ := strconv.Atoi(strconv.FormatInt(z.Milliseconds(), 10))
	fmt.Println(result)
	return result
}
