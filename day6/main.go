package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile("(\\d)+")

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	s.Scan()
	timeDigits := re.FindAllString(s.Text(), -1)
	ts := ""
	for i := range timeDigits {
		ts += timeDigits[i]
	}

	s.Scan()
	distDigits := re.FindAllString(s.Text(), -1)
	ds := ""
	for i := range distDigits {
		ds += distDigits[i]
	}
	time, _ := strconv.Atoi(ts)
	dist, _ := strconv.Atoi(ds)

	fmt.Println(ts, ds)

	res := 0
	for i := 1; i < time; i++ {
		if i*(time-i) > dist {
			res++
		}
	}
	fmt.Printf("%d\n", res)
}
