package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	res := 0

	for s.Scan() {
		line := s.Text()
		data := strings.Split(line, ":")
		games := strings.Split(data[1], ";")

		gameStats := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, g := range games {

			cubes := strings.Split(g, ",")

			for _, cube := range cubes {
				d := strings.Split(strings.TrimSpace(cube), " ")
				val, err := strconv.Atoi(d[0])
				if err != nil {
					log.Fatal(err)
				}
				col := d[1]
				if val > gameStats[col] {
					gameStats[col] = val
				}
			}
		}
		fmt.Println(line, gameStats)
		pow := gameStats["red"] * gameStats["blue"] * gameStats["green"]
		res += pow

	}
	fmt.Println(res)
}
