package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var re = regexp.MustCompile("[^\\d\\.]")

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	res := 0
	engine := make([]string, 0, 100)

	for s.Scan() {
		engine = append(engine, s.Text())
	}

	for y, line := range engine {
		fmt.Println(line)
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				start := i
				end := start
				for end < len(line) && unicode.IsDigit(rune(line[end])) {
					end++
				}
				val, err := strconv.Atoi(line[start:end])
				fmt.Println(val)
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

				var arround string
				if y > 0 {
					arround += engine[y-1][min:max]
				}
				arround += engine[y][min:max]
				if y < len(engine)-1 {
					arround += engine[y+1][min:max]
				}

				if re.MatchString(arround) {
					res += val
					fmt.Println(arround, len(arround))
				}
				i = end - 1
				continue
			}
		}
	}
	fmt.Println(res)
}
