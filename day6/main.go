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
	time := re.FindAllString(s.Text(), -1)

	s.Scan()
	dist := re.FindAllString(s.Text(), -1)

	fmt.Println(time, dist)
	vals := make([]int, 0, len(time))

	for i, t := range time {
		res := 0
		t, _ := strconv.Atoi(t)
		d, _ := strconv.Atoi(dist[i])
		for s := 1; s < t; s++ {
			if s*(t-s) > d {
				res++
			}
		}
		vals = append(vals, res)
	}
	fmt.Printf("%#v\n", vals)

	res := 0
	for _, v := range vals {
		if res == 0 {
			res = v
		} else {
			res *= v
		}
	}
	fmt.Printf("%d\n", res)
}
