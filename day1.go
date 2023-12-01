package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day1-input.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	num := []int{}

	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for s.Scan() {
		orig := s.Text()
		line := orig

		lineNums := []string{}

	first:
		for i, c := range line {
			stringChar := string(c)
			_, err := strconv.Atoi(stringChar)
			if err == nil {
				lineNums = append(lineNums, stringChar)
				break
			}

			t := line[:i+1]
			for key, val := range digits {
				if strings.Contains(t, key) {
					lineNums = append(lineNums, strconv.Itoa(val))
					fmt.Printf("Found %v in string %v \n", val, t)
					break first
				}
			}
		}

	second:
		for i := len(line) - 1; i >= 0; i-- {
			stringChar := string(line[i])
			_, err := strconv.Atoi(stringChar)
			if err == nil {
				lineNums = append(lineNums, stringChar)
				break
			}
			t := line[i:]
			for key, val := range digits {
				if strings.Contains(t, key) {
					lineNums = append(lineNums, strconv.Itoa(val))
					break second
				}
			}
		}

		n, err := strconv.Atoi(strings.Join(lineNums, ""))
		if err != nil {
			fmt.Printf("Error parsing strings %v \n", err)
		}
		fmt.Printf("Line %s, vals: %v, nums: %v \n", line, n, lineNums)
		num = append(num, n)
	}

	res := 0
	for _, v := range num {
		res += v
	}
	fmt.Println(res)
}
