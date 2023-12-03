package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var re = regexp.MustCompile("\\*")

func makeKey(y, min int) string {
	arr := []string{strconv.Itoa(y), strconv.Itoa(min)}
	return strings.Join(arr, ";")
}

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	res := 0
	engine := make([]string, 0, 100)

	for s.Scan() {
		engine = append(engine, s.Text())
	}

	symbols := map[string][]int{}

	checkString := func(y, min, max, val int) bool {
		if loc := re.FindStringIndex(engine[y][min:max]); loc != nil {
			key := makeKey(y-1, min+loc[0])
			symbols[key] = append(symbols[key], val)
			return true
		}
		return false
	}

	for y, line := range engine {
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				start := i
				end := start
				for end < len(line) && unicode.IsDigit(rune(line[end])) {
					end++
				}
				val, err := strconv.Atoi(line[start:end])
				if err != nil {
					log.Fatal(err)
				}
				min := start
				if start > 0 {
					min--
				}

				max := end
				if max < len(line) {
					max++
				}

				if y > 0 && checkString(y-1, min, max, val) {
					i = end - 1
					continue
				}
				if checkString(y, min, max, val) {
					i = end - 1
					continue
				}
				if y < len(engine)-1 && checkString(y+1, min, max, val) {
					i = end - 1
					continue
				}

				i = end - 1
				continue
			}
		}
	}
	for _, vals := range symbols {
		if len(vals) == 2 {
			res += vals[0] * vals[1]
		}
	}
	fmt.Println(symbols)
	fmt.Println(res)
}
