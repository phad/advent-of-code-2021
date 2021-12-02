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
	horiz, depth := 0, 0
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
		case "up":
			depth -= val
		case "down":
			depth += val
		}
	//	fmt.Printf("[%v, %v] %d %d\n", tokens[0], val, horiz, depth)
	}
	fmt.Printf("Final horiz / depth / product: %d %d %d\n", horiz, depth, horiz*depth)
}
