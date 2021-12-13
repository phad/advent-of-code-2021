package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateCost(pos []int, cheapest int) int {
	cost := 0
	for _, p := range pos {
		c := cheapest - p
		if c < 0 { c = -c }
		cost += c
	}
	return cost
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <in file>")
	}
	input, err := readLines(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Read %d input lines", len(input))

	if len(input) > 1 {
		log.Fatalf("too much input: got %d lines, want 1", len(input))
	}

	fmt.Println("input = "+input[0]);
	posStrs := strings.Split(input[0], ",")
	fmt.Sprintf("%v", posStrs)
	var positions []int
	min, max := 999999999, -999999999
	for _, pStr := range posStrs {
		p, err := strconv.Atoi(pStr)
		if err != nil {
			log.Fatalf("parseInput err: %v", err)
		}
		positions = append(positions, p)
		if p < min { min = p }
		if p > max { max = p }
	}

	costs := map[int]int{}
	lowest, alignToLowest := 999999999, 0
	for alignTo := min; alignTo <= max; alignTo++ {
		fmt.Printf("---> Aligning at position %d Lowest cost so far: %d --->\n", alignTo, lowest)
		cost := calculateCost(positions, alignTo)
		costs[alignTo] = cost
		if cost < lowest {
			lowest = cost
			alignToLowest = alignTo
		}
	}

	fmt.Printf("Found lowest cost of %d when aligned at %d\n", lowest, alignToLowest)
}

