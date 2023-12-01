package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type LineNums struct {
	first  string
	second string
}

func (l *LineNums) toInt() int {
	n, err := strconv.Atoi(strings.Join([]string{l.first, l.second}, ""))
	if err != nil {
		fmt.Printf("Error parsing strings %v \n", err)
	}
	return n
}

var digits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var re = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)

func checkStrings(first byte, second string) string {
	stringChar := string(first)
	_, err := strconv.Atoi(stringChar)
	if err == nil {
		return stringChar
	}

	matched := re.FindString(second)
	if matched != "" {
		return digits[matched]
	}
	return ""
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
			if lineNums.first != "" && lineNums.second != "" {
				break
			}

			if lineNums.first == "" {
				res := checkStrings(line[f], line[:f+1])
				if res != "" {
					lineNums.first = res
				}

				f++
			}
			if lineNums.second == "" {
				res := checkStrings(line[s], line[s:])
				if res != "" {
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
