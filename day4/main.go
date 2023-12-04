package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Card struct {
	points int
	count  int
}

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	res := 0
	cards := make([]Card, 0, 100)

	for s.Scan() {
		card := strings.Split(s.Text(), ":")[1]
		nums := strings.Split(card, "|")
		winning := strings.Split(strings.TrimSpace(nums[0]), " ")
		have := strings.Split(strings.TrimSpace(nums[1]), " ")

		points := 0
	o:
		for _, num := range winning {
			n := strings.TrimSpace(num)
			if n == "" {
				continue
			}
			for _, h := range have {
				if n == h {
					points++
					continue o
				}
			}
		}

		cards = append(cards, Card{points: points, count: 1})
	}

	for i, card := range cards {
		if card.points == 0 {
			continue
		}

		for j := 1; j <= card.points && i+j < len(cards); j++ {
			cards[i+j].count += card.count
		}
	}
	for _, card := range cards {
		res += card.count
	}
	fmt.Println(cards)
	fmt.Println(res)
}
