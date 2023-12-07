package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
	res := make([]int, 0, 4)
	s := strings.Split(strings.TrimSpace(line), " ")
	for _, num := range s {
		numval, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, numval)
	}
	return res
}

func parseMap(s *bufio.Scanner) [][]int {
	res := make([][]int, 0, 2)

	for s.Scan() {
		l := s.Text()

		if l == "" {
			break
		}

		if string(l[len(l)-1]) == ":" {
			continue
		}

		row := parseLine(l)
		res = append(res, row)
	}

	return res
}

func convertValue(val int, valmap [][]int) int {
	for _, row := range valmap {
		dest, src, rlen := row[0], row[1], row[2]
		if src <= val && val < (src+rlen) {
			dif := dest - src
			return val + dif
		}
	}

	return val
}

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	var seeds []int
	var soilmap, fertmap, watermap, lightmap, tempmap, humidmap, locmap [][]int

	// parse almanac
	for s.Scan() {
		l := s.Text()
		if len(seeds) == 0 {
			seedString := strings.Split(l, ": ")[1]
			seeds = parseLine(seedString)
		}

		switch true {
		case l == "seed-to-soil map:":
			soilmap = parseMap(s)
		case l == "soil-to-fertilizer map:":
			fertmap = parseMap(s)
		case l == "fertilizer-to-water map:":
			watermap = parseMap(s)
		case l == "water-to-light map:":
			lightmap = parseMap(s)
		case l == "light-to-temperature map:":
			tempmap = parseMap(s)
		case l == "temperature-to-humidity map:":
			humidmap = parseMap(s)
		case l == "humidity-to-location map:":
			locmap = parseMap(s)
		}
	}

	seedToLocation := func(seed int) int {
		soil := convertValue(seed, soilmap)
		fert := convertValue(soil, fertmap)
		water := convertValue(fert, watermap)
		light := convertValue(water, lightmap)
		temp := convertValue(light, tempmap)
		hum := convertValue(temp, humidmap)
		loc := convertValue(hum, locmap)
		return loc
	}

	res := -1
	for _, seed := range seeds {
		r := seedToLocation(seed)
		if res == -1 {
			res = r
		}
		if r < res {
			res = r
		}
	}

	fmt.Println(res)
}
