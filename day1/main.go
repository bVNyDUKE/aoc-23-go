package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"five":  5,
	"5":     5,
	"six":   6,
	"6":     6,
	"seven": 7,
	"7":     7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
}

var re = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9)`)

func checkStrings(s string) int {
	matched := re.FindString(s)
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
				res := checkStrings(line[:f+1])
				if res != 0 {
					lineNums.first = res
				}

				f++
			}
			if lineNums.second == 0 {
				res := checkStrings(line[s:])
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
