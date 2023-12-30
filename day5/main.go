package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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
		if row[1] <= val && val < (row[1]+row[2]) {
			return val + (row[0] - row[1])
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
		var i int
		i = convertValue(seed, soilmap)
		i = convertValue(i, fertmap)
		i = convertValue(i, watermap)
		i = convertValue(i, lightmap)
		i = convertValue(i, tempmap)
		i = convertValue(i, humidmap)
		i = convertValue(i, locmap)
		return i
	}

	resChan := make(chan int)
	var wg sync.WaitGroup

	crunch := func(start, end int) {
		defer wg.Done()
		res := -1
		fmt.Println("Crunching", start, end)
		for k := start; k <= end; k++ {
			val := seedToLocation(k)
			if res == -1 || val < res {
				res = val
			}
		}
		fmt.Println("Sending", res)
		resChan <- res
	}

	for i := 0; i < len(seeds)-1; i += 2 {
		start := seeds[i]
		end := seeds[i+1] + seeds[i]
		mid := (start + end) / 2
		f := (start + mid) / 2
		s := (mid + end) / 2

		wg.Add(4)
		go crunch(start, f)
		go crunch(f+1, mid)
		go crunch(mid+1, s)
		go crunch(s+1, end)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	res := -1
	for val := range resChan {
		if res == -1 || val < res {
			res = val
		}
	}

	fmt.Println("RESULT:", res)
}
