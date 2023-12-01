package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type LineNums struct {
	first  int
	second int
}

func (l *LineNums) toInt() int {
	return (l.first * 10) + l.second
}

var digits = map[string]int{
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

var re = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)

func checkStrings(first byte, second string) int {
	stringChar := string(first)
	val, err := strconv.Atoi(stringChar)
	if err == nil {
		return val
	}

	matched := re.FindString(second)
	if matched != "" {
		return digits[matched]
	}
	return 0
}

func main() {
	f, _ := os.Open("./day1-input.txt")
	s := bufio.NewScanner(f)

	res := 0

	for s.Scan() {
		line := s.Text()

		lineNums := LineNums{}
		f, s := 0, len(line)-1
		for {
			if lineNums.first != 0 && lineNums.second != 0 {
				break
			}

			if lineNums.first == 0 {
				res := checkStrings(line[f], line[:f+1])
				if res != 0 {
					lineNums.first = res
				}

				f++
			}
			if lineNums.second == 0 {
				res := checkStrings(line[s], line[s:])
				if res != 0 {
					lineNums.second = res
				}
				s--
			}

		}

		fmt.Printf("Line %s, nums: %v \n", line, lineNums)
		res += lineNums.toInt()
	}
	fmt.Println(res)
}
