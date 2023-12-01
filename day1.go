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
		line := s.Text()

		lineNums := struct {
			first  string
			second string
		}{
			first:  "",
			second: "",
		}
		f, s := 0, len(line)-1
		for {
			if lineNums.first != "" && lineNums.second != "" {
				break
			}

			if lineNums.first == "" {
				stringChar := string(line[f])
				_, err := strconv.Atoi(stringChar)
				if err == nil {
					lineNums.first = stringChar
				}

				t := line[:f+1]
				for key, val := range digits {
					if strings.Contains(t, key) {
						lineNums.first = strconv.Itoa(val)
						fmt.Printf("Found %v in string %v \n", val, t)
						break
					}
				}
				f++
			}
			if lineNums.second == "" {
				stringChar := string(line[s])
				_, err := strconv.Atoi(stringChar)
				if err == nil {
					lineNums.second = stringChar
				}

				t := line[s:]
				for key, val := range digits {
					if strings.Contains(t, key) {
						lineNums.second = strconv.Itoa(val)
						break
					}
				}
				s--
			}

		}

		n, err := strconv.Atoi(strings.Join([]string{lineNums.first, lineNums.second}, ""))
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
