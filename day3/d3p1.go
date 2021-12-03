package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <in file>")
	}
	lines, err := readLines(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Read %d input lines", len(lines))

	type counts struct{ n0, n1 int }
	var bitCounts []counts

	for _, bin := range lines {
		if bin == "" {
			continue
		}
		sz := len(bin)
		for len(bitCounts) < sz {
			bitCounts = append(bitCounts, counts{})
		}
		for pos := sz-1; pos >= 0; pos-- {
			if bin[pos] == '0' {
				bitCounts[pos].n0++;
			} else if bin[pos] == '1' {
				bitCounts[pos].n1++;
			}
		}
	}

	var gamma, epsilon int
	for pos, m := len(bitCounts)-1, 1; pos >= 0; pos-- {
		if bitCounts[pos].n1 > bitCounts[pos].n0 {
			gamma += m
		}
		if bitCounts[pos].n0 > bitCounts[pos].n1 {
			epsilon += m
		}
		m <<= 1
	}

	fmt.Printf("gamma = %d, epsilon = %d, product: %d\n", gamma, epsilon, gamma*epsilon)
}
