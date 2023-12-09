package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getUniqueChars(s string) int {
	u := map[string]int{}
	jokerCount := 0
	for _, char := range strings.Split(s, "") {
		if char == "J" {
			jokerCount++
		} else {
			u[char]++
		}
	}

	vals := make([]int, 0, 5)
	for _, val := range u {
		vals = append(vals, val)
	}
	slices.Sort(vals)
	slices.Reverse(vals)

	res := make([]int, 5, 5)
	for i, val := range vals {
		res[i] = val
	}

	res[0] = res[0] + jokerCount

	rank := 0
	for i, val := range res {
		rank += (val * (int(math.Pow10(4 - i))))
	}

	return rank
}

var charVals = map[string]int{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": -1,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	cards := make([][]string, 0, 5)

	for s.Scan() {
		card := strings.Split(s.Text(), " ")
		cards = append(cards, card)
	}

	slices.SortFunc(cards, func(first, second []string) int {
		firstCard, secondCard := first[0], second[0]

		fu, su := getUniqueChars(firstCard), getUniqueChars(secondCard)

		if fu > su {
			return 1
		}
		if fu < su {
			return -1
		}

		val := 0
		for i := 0; i < 5; i++ {
			ff, sf := string(firstCard[i]), string(secondCard[i])
			// fmt.Println(firstCard, secondCard, ff, sf)
			if charVals[ff] < charVals[sf] {
				val = -1
				break
			}
			if charVals[ff] > charVals[sf] {
				val = 1
				break
			}
		}
		return val
	})

	res := 0
	for i, card := range cards {
		fmt.Println(card)
		h, err := strconv.Atoi(card[1])
		if err != nil {
			log.Fatal(err)
		}
		res += h * (i + 1)
	}

	fmt.Printf("%#v \n", res)
}
