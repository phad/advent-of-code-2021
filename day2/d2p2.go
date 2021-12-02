package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	horiz, depth, aim := 0, 0, 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		tokens := strings.Split(l, " ")
		val, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("Atoi: %v", err)
		}
		switch tokens[0] {
		case "forward":
			horiz += val
			depth += aim*val
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	//	fmt.Printf("[%v, %v] %d %d %d\n", tokens[0], val, horiz, depth, aim)
	}
	fmt.Printf("Final horiz / depth / product: %d %d %d\n", horiz, depth, horiz*depth)
}
