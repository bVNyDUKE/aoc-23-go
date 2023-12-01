package main

import (
	"bufio"
	"fmt"
	"os"
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

func checkStrings(first byte, second string) string {
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
	stringChar := string(first)
	_, err := strconv.Atoi(stringChar)
	if err == nil {
		return stringChar
	}

	for key, val := range digits {
		if strings.Contains(second, key) {
			return strconv.Itoa(val)
		}
	}
	return ""
}

func main() {
	f, _ := os.Open("./day1-input.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	num := []int{}

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
		num = append(num, lineNums.toInt())
	}

	res := 0
	for _, v := range num {
		res += v
	}
	fmt.Println(res)
}
