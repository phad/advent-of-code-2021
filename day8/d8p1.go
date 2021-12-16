package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type scrambling struct {
	digits []string
	output []string
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

	// each input line:
	// 10* scrambled segment states | 4* output digits
	// scrambling is consistent for the whole line, but
	// differs from one line to another.

	// abcdefg refer to 7 segments.  We know that digits
	// 1: lights two segments
	// 7: lights three segments
	// 4: lights four segments
	// 2, 3, 5: light five segments
	// 0, 6, 9: light six segments
	// 7: lights seven segments

	var scramblings []*scrambling

	for _, in := range input {
		fmt.Println(in)
		parts := strings.Split(in, " | ")
		if len(parts) == 0 {
			continue
		}
		if len(parts) != 2 {
			log.Fatalf("Malformed input: %q", in)
		}
		digits := strings.Split(parts[0], " ")
		if len(digits) != 10 {
			log.Fatalf("Malformed digits: %q", parts[0])
		}
		output := strings.Split(parts[1], " ")
		if len(output) != 4 {
			log.Fatalf("Malformed output: %q", parts[1])
		}
		scr := &scrambling{
			digits: digits,
			output: output,
		}
		fmt.Printf("Read scrambling: %v\n", scr)
		scramblings = append(scramblings, scr)
	}

	// Part 1: how many times do digits 1, 4, 7, or 8 appear.
	count1478 := 0
	for _, scr := range scramblings {
		for _, outDigit := range scr.output {
			switch len(outDigit) {
			case 2, 3, 4, 7:
				count1478++
			default:
			}
		}
	}
	fmt.Printf("Digits 1, 4, 7 or 8 appear %d times in input.\n", count1478)
}

