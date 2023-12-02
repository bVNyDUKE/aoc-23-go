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

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for s.Scan() {
		line := s.Text()
		data := strings.Split(line, ":")
		gameId, _ := strconv.Atoi(strings.Split(data[0], " ")[1])
		games := strings.Split(data[1], ";")

		valid := true
	o:
		for _, g := range games {
			cubes := strings.Split(g, ",")
			for _, cube := range cubes {
				d := strings.Split(strings.TrimSpace(cube), " ")
				val, err := strconv.Atoi(d[0])
				if err != nil {
					log.Fatal(err)
				}
				col := d[1]
				if val > limits[col] {
					valid = false
					break o
				}
			}
		}
		if valid {
			fmt.Println("Possible game", line)
			res += gameId
		}

	}
	fmt.Println(res)
}
