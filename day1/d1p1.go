package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	first := true
	prev := 0
	numInc, numDec, numSame := 0, 0, 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		val, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("Atoi: %v", err)
		}
		if !first {
			switch {
			case val < prev:
				numDec++
			case val > prev:
				numInc++
			default:
				numSame++
			}
//			fmt.Printf("[%v, %v] %d %d %d\n", prev, val, numDec, numSame, numInc)
		}
		first = false
		prev = val
	}
	fmt.Printf("Num decreased / same / increased: %d %d %d\n", numDec, numSame, numInc)
}
