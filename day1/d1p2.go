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

	var window [3]int
	prev := 0
	numInc, numDec, numSame := 0, 0, 0
	for idx, l := range lines {
		if l == "" {
			continue
		}
		val, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("Atoi: %v", err)
		}
		window[idx%3] = val
		sum := window[0]+window[1]+window[2]

		//fmt.Printf("%v %d: ", window, sum)
		if idx < 3 {
			//fmt.Print("\n")
			prev = sum
			continue
		}
		switch {
		case sum < prev:
			numDec++
		case sum > prev:
			numInc++
		default:
			numSame++
		}
		//fmt.Printf("[%v, %v] %d %d %d\n", prev, sum, numDec, numSame, numInc)
		prev = sum
	}
	fmt.Printf("Num decreased / same / increased: %d %d %d\n", numDec, numSame, numInc)
}
